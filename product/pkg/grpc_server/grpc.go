package grpc_server

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	Port    string
	Options []grpc.ServerOption
}

func New(port string, opts ...grpc.ServerOption) *GRPCServer {
	return &GRPCServer{
		Port:    port,
		Options: opts,
	}
}

// Start Принимает список функций-регистраторов, например: pb.RegisterProductServiceServer
func (s *GRPCServer) Start(registrations ...func(*grpc.Server)) error {
	lis, err := net.Listen("tcp", ":"+s.Port)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer(s.Options...)
	for _, register := range registrations {
		register(grpcServer)
	}

	log.Printf("🚀 gRPC server started on :%s\n", s.Port)
	return grpcServer.Serve(lis)
}
