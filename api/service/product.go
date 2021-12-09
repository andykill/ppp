package service

import (
	context "context"
	"github.com/golang/protobuf/ptypes/empty"
	"ppp/entity"
	"ppp/pb"
)

type Product struct {
	warehouse *Warehouse
	user      *entity.User
	pb.UnsafeProductServiceServer
}

func NewProductService(warehouse *Warehouse, user *entity.User) *Product {
	return &Product{warehouse: warehouse, user: user}
}

func (p Product) ProductChangeCount(ctx context.Context, product *pb.Product) (*empty.Empty, error) {

	err := p.warehouse.ChangeCountProduct(product.Id, product.Count)

	return nil, err
}
