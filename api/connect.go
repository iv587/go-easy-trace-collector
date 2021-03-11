package api

import (
	"collector/client"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type connect struct {
}

var pConnect = new(connect)

func (t *connect) list(c *gin.Context) {
	conn := client.ConnectList()
	list := make([]ConnectInfoVo, 0, len(conn))
	for k, v := range conn {
		cVo := ConnectInfoVo{
			Addr:         k,
			AppName:      v.AppInfo.Name,
			AppGroup:     v.AppInfo.Group,
			CreateTime:   v.Time.Format("2006-01-02 15:04:05"),
			AppStartTime: time.Unix(v.AppInfo.StartTime/1000, (v.AppInfo.StartTime%1000)*1000000).Format("2006-01-02 15:04:05"),
			AliveTime:    parseMillToStr(v.Alive.Sub(v.Time).Milliseconds()),
			NetInput:     translateBytes(v.ReadBytes.Gb, v.ReadBytes.Mb, v.ReadBytes.Kb, v.ReadBytes.By),
		}
		list = append(list, cVo)
	}
	succ(c, "", gin.H{
		"list": list,
	})
}

func parseMillToStr(milliSec int64) string {
	// ms
	if milliSec < 1000 {
		return fmt.Sprintf("%dms", milliSec)
	}
	// s
	sec := milliSec / 1000
	if sec < 60 {
		return fmt.Sprintf("%ds", sec)
	}
	// min
	if sec < 3600 {
		min := sec / 60
		str := fmt.Sprintf("%dmin", min)
		mod := sec % 60
		if mod > 0 {
			str = str + fmt.Sprintf("%ds", mod)
		}
		return str
	}
	// hour
	if sec < 86400 {
		min := sec / 3600
		str := fmt.Sprintf("%dh", min)
		mod := (sec % 3600) / 60
		if mod > 0 {
			str = str + fmt.Sprintf("%dmin", mod)
		}
		return str
	}
	// day
	day := sec / 86400
	str := fmt.Sprintf("%dday", day)
	mod := (sec % 86400) / 3600
	if mod > 0 {
		str = str + fmt.Sprintf("%dh", mod)
	}
	return str

}

func translateBytes(GB, MB, KB, By int) string {
	if GB > 1 {
		d := (MB % 1000) / 100
		if d > 0 {
			return fmt.Sprintf("%d.%dG", GB, d)
		} else {
			return fmt.Sprintf("%dG", GB)
		}
	}
	if MB > 1 {
		d := (KB % 1000) / 100
		if d > 0 {
			return fmt.Sprintf("%d.%dM", MB, d)
		} else {
			return fmt.Sprintf("%dM", MB)
		}
	}
	if KB > 1 {
		d := (By % 1000) / 100
		if d > 0 {
			return fmt.Sprintf("%d.%dK", KB, d)
		} else {
			return fmt.Sprintf("%dK", KB)
		}
	}
	return fmt.Sprintf("%dbytes", By)

}
