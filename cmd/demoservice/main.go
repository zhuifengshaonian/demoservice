package main

import (
	"demoservice/configs"
	"demoservice/internal/pkg/zlog"
	"demoservice/internal/server"
	"flag"
)

var flagConf string

func init() {
	flag.StringVar(&flagConf, "conf", "../../configs", "config path, eg: -conf config.yaml")

}

func main() {
	flag.Parse()
	appInfo := configs.Load(&flagConf)

	zlog.Init(appInfo.Log.FilePath)
	defer zlog.Sync()

	zlog.Sugar.Info(appInfo)
	err := server.Server(appInfo.Grpc, appInfo.Http)
	if err != nil {
		panic("start server failed.")
	}

}
