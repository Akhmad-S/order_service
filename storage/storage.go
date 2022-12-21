package storage

import ecom "github.com/uacademy/e_commerce/order_service/proto-gen/e_commerce"

type StorageI interface{
	CreateOrder(id string, input *ecom.CreateOrderRequest) error
	GetOrderList(offset, limit int, search string) (resp *ecom.GetOrderListResponse, err error)
	GetOrderById(id string) (*ecom.GetOrderByIdResponse, error)
}
