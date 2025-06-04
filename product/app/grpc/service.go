package grpc

import productv1 "github.com/Deevins/lampshop-backend/product/gen/github.com/Deevins/lampshop-backend/product/v1"

type Implementation struct {
	productv1.UnimplementedProductServiceServer

	// TODO: db adapter
}

func NewProductAPI() *Implementation {
	return &Implementation{}
}
