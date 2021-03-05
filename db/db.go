package db

import (
	"collector/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func Boot() error {
	var err error
	engine, err = xorm.NewEngine("mysql", config.Mysql.Addr)
	if err != nil {
		return err
	}
	engine.SetMaxIdleConns(3)
	engine.SetMaxOpenConns(20)
	return nil
}

func GetEngine() *xorm.Engine {
	return engine
}
