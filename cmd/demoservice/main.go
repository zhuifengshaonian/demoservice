package main

import (
	"demoservice/internal/pkg/zlog"
	"demoservice/internal/server"
	"flag"
	"fmt"

	"github.com/spf13/viper"
)

var flagConf string

func init() {
	flag.StringVar(&flagConf, "conf", "../../configs", "config path, eg: -conf config.yaml")

}

func main() {
	flag.Parse()
	viper.SetConfigFile(flagConf)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("load config file %v failed.", flagConf))
	}

	zlog.Init(viper.GetString("log.file-path"))
	defer zlog.Sync()
	zlog.Sugar.Infoln(viper.AllSettings())
	err = server.Server()
	if err != nil {
		panic("start server failed.")
	}

}
