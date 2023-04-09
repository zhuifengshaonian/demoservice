package server

import (
	"context"
	v1 "demoservice/api/demoservice/v1"
	"demoservice/internal/pkg/zlog"
	"demoservice/internal/service"
	"net"
	"net/http"

	"google.golang.org/grpc/reflection"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Server() error {
	addr := ":9000"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	v1.RegisterDemoServiceServer(s, service.NewDemoService())
	reflection.Register(s)
	// Serve gRPC server
	go func() {
		zlog.Sugar.Fatalln(s.Serve(lis))
	}()
	zlog.Sugar.Infof("serving grpc on %s", addr)

	//创建一个已经启动的grpc server的客户端连接
	conn, err := grpc.DialContext(
		context.Background(),
		addr,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		zlog.Sugar.Infof("dial failed:%v\n", err)
		return err
	}
	gwMux := runtime.NewServeMux()
	err = v1.RegisterDemoServiceHandler(context.Background(), gwMux, conn)
	if err != nil {
		zlog.Sugar.Infof("register handler failed:%v\n", err)
		return err
	}
	server := http.Server{
		Addr:    ":9001",
		Handler: gwMux,
	}
	zlog.Sugar.Infof("serving gateway on %s", server.Addr)
	zlog.Sugar.Fatalln(server.ListenAndServe())
	return nil
}
