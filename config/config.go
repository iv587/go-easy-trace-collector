package config

import (
	"github.com/spf13/viper"
	"os"
)

type config struct {
}

var Mysql = struct {
	Addr string
}{}

var Http = struct {
	Addr string
}{}

var Collector = struct {
	Addr string
}{}

var User = struct {
	Name     string
	Password string
}{}

func Load() error {
	viper.SetConfigName("app")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	Mysql.Addr = viper.GetString("mysql.addr")
	Http.Addr = viper.GetString("http.addr")
	Collector.Addr = viper.GetString("collector.addr")
	User.Name = viper.GetString("user.name")
	User.Password = viper.GetString("user.password")
	val := os.Getenv("MYSQL_ADDR")
	if val != "" {
		Mysql.Addr = val
	}
	val = os.Getenv("COLLECTOR_ADDR")
	if val != "" {
		Collector.Addr = val
	}
	val = os.Getenv("HTTP_ADDR")
	if val != "" {
		Http.Addr = val
	}
	val = os.Getenv("USER_NAME")
	if val != "" {
		User.Name = val
	}
	val = os.Getenv("USER_PASSWORD")
	if val != "" {
		User.Password = val
	}
	return nil

}
