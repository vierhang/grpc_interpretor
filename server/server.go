package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_interpretor/proto"
	"net"
)

type Server struct {
}

func (s Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{
		Message: request.Name,
	}, nil
}

func main() {
	// 自定义拦截器逻辑
	myInterceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("接收到了一个请求")
		res, err := handler(ctx, req)
		fmt.Println("请求完成")
		return res, err
	}
	// 拦截器
	opt := grpc.UnaryInterceptor(myInterceptor)
	g := grpc.NewServer(opt)
	proto.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		panic(err)
	}
	err = g.Serve(lis)
	if err != nil {
		panic(err)
	}
}
