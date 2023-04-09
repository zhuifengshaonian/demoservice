package configs

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
)

var v *viper.Viper

type AppInfo struct {
	Service          string      `yaml:"service"`
	Group            string      `yaml:"group"`
	MicroServiceName string      `yaml:"micro-service-name" mapstructure:"micro-service-name"`
	Log              *LogConfig  `yaml:"log"`
	Http             *HttpConfig `yaml:"http"`
	Grpc             *GrpcConfig `yaml:"grpc"`
}

type LogConfig struct {
	FilePath string `yaml:"file-path" mapstructure:"file-path"`
}

type GrpcConfig struct {
	Addr string `yaml:"addr"`
	Port int32  `yaml:"port"`
}

type HttpConfig struct {
	Addr string `yaml:"addr"`
	Port int32  `yaml:"port"`
}

func Load(filePath *string) (app *AppInfo) {
	v = viper.New()
	v.SetConfigFile(*filePath)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("load config file %v failed.", filePath))
	}
	app = &AppInfo{}
	err = v.Unmarshal(app)
	if err != nil {
		return
	}
	return app
}

func (AppInfo) String() string {
	marshal, err := json.MarshalIndent(v.AllSettings(), "", "\t")
	if err != nil {
		return fmt.Sprintf("%v", v.AllSettings())
	}
	return string(marshal)
}
