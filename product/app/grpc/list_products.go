package grpc

import (
	"context"
	desc "github.com/Deevins/lampshop-backend/product/gen/github.com/Deevins/lampshop-backend/product/v1"
)

func (i *Implementation) ListProducts(ctx context.Context, req *desc.ListProductsRequest) (*desc.ListProductsResponse, error) {
	return &desc.ListProductsResponse{
		Products: []*desc.Product{
			{
				Id:   "123",
				Name: "product1",
			},
		},
	}, nil
}
