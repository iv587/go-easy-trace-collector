package client

import (
	"bufio"
	"bytes"
	"collector/easytrace"
	"collector/span"
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"net"
	"sync"
	"time"
)

const (
	timeMaxOut = time.Second * 30
)

var pongBytes = []byte{0, 1}

type Context struct {
	// 原生链接
	conn net.Conn
	// 链接时间
	Time time.Time
	// 存活时间
	Alive time.Time
	// 事件
	running bool

	ReadBytes struct {
		sync.RWMutex
		Gb int
		Mb int
		Kb int
		By int
	}
	AppInfo AppInfo
}

// 创建链接上下文
func Create(conn net.Conn) {
	ctx := &Context{
		conn:    conn,
		Time:    time.Now(),
		Alive:   time.Now(),
		running: true,
	}
	ctx.start()
	registContext(ctx)
}

func (c *Context) start() {
	// 启动消息接收
	go c.recvData()
	go c.aliveCheck()
}

func (c *Context) Write(msg []byte) (int, error) {
	n, err := c.conn.Write(msg)
	if err == nil {
		c.Alive = time.Now()
	}
	return n, err
}

func (c *Context) close() {
	c.conn.Close()
	c.running = false
	removeContext(c)
}

func handleRecv(c *Context, packet Packet) error {
	if packet.Type == typeHeartbeat {
		var heartBeat easytrace.HeartBeatReq
		proto.Unmarshal(packet.Body, &heartBeat)
		c.AppInfo.Name = heartBeat.AppName
		c.AppInfo.Group = heartBeat.GroupName
		c.AppInfo.StartTime = heartBeat.TimeStamp
		fmt.Println(heartBeat)
	} else {
		body := make([]byte, len(packet.Body))
		copy(body, packet.Body)
		go span.Process(body)
	}
	return nil
}

func (c *Context) recvData() {
	defer c.close()
	r := bufio.NewReader(c.conn)
	body := make([]byte, 2*1024*1024)
	bodyLen := 0
	var msgType int32 = 0
	cache := false
	bBuff := bytes.NewBuffer(nil)
	cBuff := bytes.NewBuffer(nil)
	msgBodyList := make([]Packet, 0, 10)
	for c.running {
		n, err := r.Read(body)
		if err != nil {
			return
		}
		c.Alive = time.Now()
		if n > 0 {
			fmt.Println("接受", n)
			// 记录流量
			c.countReadBytes(n)
			msgBodyList = msgBodyList[0:0]
			bBuff.Reset()
			if cache {
				bBuff.Write(cBuff.Bytes())
			}
			bBuff.Write(body[:n])
			cache, bodyLen, msgType = c.decode(bBuff, cBuff, bodyLen, msgType, &msgBodyList)
			// 处理消息体
			if len(msgBodyList) > 0 {
				for _, msgBody := range msgBodyList {
					err := handleRecv(c, msgBody)
					if err != nil {
						return
					}
				}
			}
		}
	}
}

func (c *Context) decode(bBuff, cBuff *bytes.Buffer, len int, oldmsgSign int32, msgBodyList *[]Packet) (bool, int, int32) {
	cBuff.Reset()
	bodyLen := len
	cache := false
	index := 0
	var msgSign int32
	for bBuff.Len() > 0 {
		if bodyLen <= 0 {
			if bBuff.Len() >= 8 {
				binary.Read(bytes.NewReader(bBuff.Next(4)), binary.BigEndian, &msgSign)
				if msgSign != typeSpan && msgSign != typeHeartbeat && msgSign != typeHeartbeatRes {
					continue
				}
				oldmsgSign = msgSign
				var lenTmp int32
				binary.Read(bytes.NewReader(bBuff.Next(4)), binary.BigEndian, &lenTmp)
				bodyLen = int(lenTmp)
			} else {
				cBuff.Write(bBuff.Next(bBuff.Len()))
				cache = true
				break
			}
		} else {
			if bBuff.Len() >= bodyLen {
				msgBody := bBuff.Next(bodyLen)
				*msgBodyList = append(*msgBodyList, Packet{Type: oldmsgSign, Body: msgBody})
				index++
				bodyLen = 0
				msgSign = 0
			} else {
				cBuff.Write(bBuff.Next(bBuff.Len()))
				cache = true
				break
			}
		}
	}
	return cache, bodyLen, msgSign
}

func (c *Context) aliveCheck() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	for c.running {
		time := <-ticker.C
		timeOut := time.Sub(c.Alive)
		if timeOut > timeMaxOut {
			c.close()
		}
	}
}

func (c *Context) countReadBytes(n int) {
	c.ReadBytes.Lock()
	defer c.ReadBytes.Unlock()
	c.ReadBytes.By = c.ReadBytes.By + n
	if c.ReadBytes.By >= 1000 {
		count, mod := convertByte(c.ReadBytes.By)
		c.ReadBytes.Kb = c.ReadBytes.Kb + count
		c.ReadBytes.By = mod
	}
	if c.ReadBytes.Kb >= 1000 {
		count, mod := convertByte(c.ReadBytes.Kb)
		c.ReadBytes.Mb = c.ReadBytes.Mb + count
		c.ReadBytes.Kb = mod
	}
	if c.ReadBytes.Mb >= 1000 {
		count, mod := convertByte(c.ReadBytes.Mb)
		c.ReadBytes.Gb = c.ReadBytes.Gb + count
		c.ReadBytes.Mb = mod
	}
}

func convertByte(bNum int) (int, int) {
	count := bNum / 1000
	mod := bNum % 1000
	return count, mod
}
