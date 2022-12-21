package clients

import (
	"github.com/uacademy/e_commerce/order_service/config"
	ecom "github.com/uacademy/e_commerce/order_service/proto-gen/e_commerce"

	"google.golang.org/grpc"
)

type GrpcClients struct {
	Product ecom.ProductServiceClient
	Category ecom.CategoryServiceClient
}

func NewGrpcClients(cfg config.Config) (*GrpcClients, error) {
	connCategory, err := grpc.Dial(cfg.CatalogServiceGrpcHost+cfg.CatalogServiceGrpcPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	category := ecom.NewCategoryServiceClient(connCategory)
	
	connProduct, err := grpc.Dial(cfg.CatalogServiceGrpcHost+cfg.CatalogServiceGrpcPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	product := ecom.NewProductServiceClient(connProduct)


	return &GrpcClients{
		Category: category,
		Product: product,
	}, nil
}
