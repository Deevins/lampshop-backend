package main

import (
	grpc2 "github.com/Deevins/lampshop-backend/product/app/grpc"
	productv1 "github.com/Deevins/lampshop-backend/product/gen/github.com/Deevins/lampshop-backend/product/v1"
	"github.com/Deevins/lampshop-backend/product/pkg/grpc_server"
	"google.golang.org/grpc"
	"log"
)

func main() {

	srv := grpc_server.New("8081")
	productAPI := grpc2.NewProductAPI()

	if err := srv.Start(
		func(s *grpc.Server) {
			productv1.RegisterProductServiceServer(s, productAPI)
		}); err != nil {
		log.Fatalf("gRPC server failed: %v", err)
	}
}
