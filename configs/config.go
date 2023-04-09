package configs

type AppInfo struct {
	Service          string
	Group            string
	MicroServiceName string
	Log              *LogConfig
	Http             *HttpConfig
	Grpc             *GrpcConfig
}

type LogConfig struct {
	FilePath string
}

func (l LogConfig) String() string {
	return ""
}

type GrpcConfig struct {
	Addr string
	Port int32
}

func (g GrpcConfig) String() string {
	return ""
}

type HttpConfig struct {
	Addr string
	Port int32
}

func (h HttpConfig) String() string {
	return ""
}

func (a AppInfo) String() string {
	return ""
}
