package service

import (
	infrastructures "kang-sayur-backend/infrastructure/injector"
	"kang-sayur-backend/model/domain/cart"
	request_body "kang-sayur-backend/model/web/request_body/cart"
)

type cartService struct {
	repository cart.CartRepository
}

// Create implements cart.CartRepository
func (*cartService) Create(data *request_body.BasicAction) (*cart.Cart, error) {
	panic("unimplemented")
}

// Delete implements cart.CartRepository
func (*cartService) Delete(data *request_body.BasicAction) error {
	panic("unimplemented")
}

// ReadCart implements cart.CartRepository
func (*cartService) ReadCart(data *request_body.ReadCart) ([]cart.Cart, error) {
	panic("unimplemented")
}

// UpdateQuantity implements cart.CartRepository
func (*cartService) UpdateQuantity(data *request_body.UpdateQuantity) error {
	panic("unimplemented")
}

func CartService(infra *infrastructures.Infrastructures) cart.CartRepository {
	return &cartService{repository: infra.Repositories().Cart}
}
