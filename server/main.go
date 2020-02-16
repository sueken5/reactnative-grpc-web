package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"

	"github.com/sueken5/reactnative-grpc-web/server/apis/proto/ping"
)

func main() {
	fmt.Println("hello")

	s := grpc.NewServer()

	ping.RegisterPingServer(s, NewServer())
	reflection.Register(s)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9090))
	if err != nil {
		panic(err)
	}

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

type Server struct {
	
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Ping(ctx context.Context, req *ping.PingRequest) (*ping.PingResponse, error) {
	return &ping.PingResponse{Message:fmt.Sprintf("%s, world", req.Message)}, nil
}