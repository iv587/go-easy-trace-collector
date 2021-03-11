package server

import (
	"collector/client"
	"collector/config"
	"net"
)

func Boot() error {
	l, err := net.Listen("tcp", config.Collector.Addr)
	if err != nil {
		return err
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}
		client.Create(conn)
	}
}
