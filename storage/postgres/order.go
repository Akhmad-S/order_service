package postgres

import (
	"errors"

	ecom "github.com/uacademy/e_commerce/order_service/proto-gen/e_commerce"
)

func (stg Postgres) CreateOrder(id string, input *ecom.CreateOrderRequest) error {
	_, err := stg.db.Exec(`INSERT INTO "order" (id, product_id, quantity, user_name, user_address, user_phone) VALUES ($1, $2, $3, $4, $5, $6)`, id, input.ProductId, input.Quantity, input.UserName, input.UserAddress, input.UserPhone)
	if err != nil {
		return err
	}
	return nil
}

func (stg Postgres) GetOrderList(offset, limit int, search string) (resp *ecom.GetOrderListResponse, err error) {
	resp = &ecom.GetOrderListResponse{
		Orders: make([]*ecom.Order, 0),
	}

	rows, err := stg.db.Queryx(`SELECT
	id,
	product_id,
	quantity,
	user_name,
	user_address,
	user_phone,
	created_at
	FROM "order" WHERE (user_address ILIKE '%' || $1 || '%')
	LIMIT $2
	OFFSET $3
	`, search, limit, offset)

	if err != nil {
		return resp, err
	}
	for rows.Next() {
		o := &ecom.Order{}

		err := rows.Scan(
			&o.Id,
			&o.ProductId,
			&o.Quantity,
			&o.UserName,
			&o.UserAddress,
			&o.UserPhone,
			&o.CreatedAt,
		)
		if err != nil {
			return resp, err
		}

		resp.Orders = append(resp.Orders, o)
	}

	return resp, err
}

func (stg Postgres) GetOrderById(id string) (*ecom.GetOrderByIdResponse, error) {
	res := &ecom.GetOrderByIdResponse{
		Product: &ecom.GetOrderByIdResponse_Product{},
	}

	err := stg.db.QueryRow(`SELECT
		id, product_id, quantity, user_name, user_address, user_phone, created_at
		FROM "order" WHERE id = $1`, id).Scan(
		&res.Id, &res.Product.Id, &res.Quantity, &res.UserName, &res.UserAddress, &res.UserPhone, &res.CreatedAt,
	)
	if err != nil {
		return res, errors.New("order not found")
	}

	return res, nil
}
