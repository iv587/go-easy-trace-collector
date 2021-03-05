package bootstrap

import (
	"collector/config"
	"collector/db"
	"collector/http"
	"collector/server"
)

func Start() error {
	// 加载配置
	err := config.Load()
	if err != nil {
		return err
	}
	// 启动数据
	err = db.Boot()
	if err != nil {
		return err
	}
	//启动http服务
	go http.Start()
	// 启动 收集器服务
	err = server.Boot()
	if err != nil {
		return err
	}
	return nil
}
