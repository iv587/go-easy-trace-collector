package server

import (
	"collector/config"
	"collector/connection"
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
		connection.Create(conn)
	}
}
