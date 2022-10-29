package service

import (
	infrastructures "kang-sayur-backend/infrastructure/injector"
	"kang-sayur-backend/model/domain/cart"
	response "kang-sayur-backend/model/web"
	request_body "kang-sayur-backend/model/web/request_body/cart"
)

type cartService struct {
	repository cart.CartRepository
}

// Add implements cart.CartService
func (cs *cartService) Add(body *request_body.BasicAction) *response.HTTPResponse {
	cart, err := cs.repository.Create(body)

	if err != nil {
		return &response.HTTPResponse{
			Message: "Kesalahan saat penyimpanan data",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "berhasil ditambahkan ke keranjang",
		Status:    200,
		IsSuccess: true,
		Data:      cart,
	}
}

// Delete implements cart.CartService
func (cs *cartService) Delete(body *request_body.BasicAction) *response.HTTPResponse {
	if err := cs.repository.Delete(body); err != nil {
		return &response.HTTPResponse{
			Message: "Kesalahan saat menghapus data",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "Barang berhasil dihapus dari keranjang",
		Status:    200,
		IsSuccess: true,
	}
}

// GetCart implements cart.CartService
func (cs *cartService) GetCart(body *request_body.ReadCart) *response.HTTPResponse {
	carts, err := cs.repository.ReadCart(body)

	if err != nil {
		return &response.HTTPResponse{
			Message: "Kesalahan saat pengambilan data",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "keranjang berhasil diambil",
		Status:    200,
		IsSuccess: true,
		Data:      carts,
	}
}

// UpdateQuantity implements cart.CartService
func (cs *cartService) UpdateQuantity(body *request_body.UpdateQuantity) *response.HTTPResponse {
	if err := cs.repository.UpdateQuantity(body); err != nil {
		return &response.HTTPResponse{
			Message: "kesalahan saat update quantitas",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "Quantitas berhasil diupdate",
		Status:    200,
		IsSuccess: true,
	}
}

func CartService(infra *infrastructures.Infrastructures) cart.CartService {
	return &cartService{repository: infra.Repositories().Cart}
}
