package service

import (
	"fmt"
	infrastructures "kang-sayur-backend/infrastructure/injector"
	"kang-sayur-backend/model/domain/grocery"
	response "kang-sayur-backend/model/web"
	request_body "kang-sayur-backend/model/web/request_body/grocery"
)

type groceryService struct {
	repository grocery.GroceryRepository
}

// Create implements grocery.GroceryService
func (gs *groceryService) Create(body *request_body.Create) *response.HTTPResponse {
	inserted, err := gs.repository.Create(body)

	if err != nil {
		return &response.HTTPResponse{
			Message: "kesalahan saat menyimpan data",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "barang berhasil disimpan",
		Status:    200,
		IsSuccess: true,
		Data:      inserted,
	}
}

// Delete implements grocery.GroceryService
func (gs *groceryService) Delete(body *request_body.UpdateOrDelete) *response.HTTPResponse {
	if err := gs.repository.Delete(body); err != nil {
		return &response.HTTPResponse{
			Message: "kesalahan saat pengapusan data",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "produk berhasil dihapus",
		Status:    200,
		IsSuccess: true,
	}
}

// Read implements grocery.GroceryService
func (gs *groceryService) Read(body *request_body.Create) *response.HTTPResponse {
	feedbacks, err := gs.repository.Read(body)

	if err != nil {
		fmt.Println(err)
		return &response.HTTPResponse{
			Message: "kesalahan saat pengambilan data",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "grocery berhasil diambil",
		Status:    200,
		IsSuccess: true,
		Data:      feedbacks,
	}
}

// ReadBestSeller implements grocery.GroceryService
func (gs *groceryService) ReadBestSeller() *response.HTTPResponse {
	feedbacks, err := gs.repository.ReadBestSeller()

	if err != nil {
		fmt.Println(err)
		return &response.HTTPResponse{
			Message: "kesalahan saat pengambilan data",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "grocery terlaku berhasil diambil",
		Status:    200,
		IsSuccess: true,
		Data:      feedbacks,
	}
}

// ReadByCategory implements grocery.GroceryService
func (gs *groceryService) ReadByCategory(body *request_body.ByCategory) *response.HTTPResponse {
	feedbacks, err := gs.repository.ReadByCategory(body)

	if err != nil {
		fmt.Println(err)
		return &response.HTTPResponse{
			Message: "kesalahan saat pengambilan data",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "produk berdasarkan kategori berhasil diambil",
		Status:    200,
		IsSuccess: true,
		Data:      feedbacks,
	}
}

// ReadByKeyword implements grocery.GroceryService
func (gs *groceryService) ReadByKeyword(body *request_body.ByKeyword) *response.HTTPResponse {
	feedbacks, err := gs.repository.ReadByKeyword(body)

	if err != nil {
		fmt.Println(err)
		return &response.HTTPResponse{
			Message: "kesalahan saat pengambilan data",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "produk berdasarkan kata kunci berhasil diambil",
		Status:    200,
		IsSuccess: true,
		Data:      feedbacks,
	}
}

// ReadCheapest implements grocery.GroceryService
func (gs *groceryService) ReadCheapest() *response.HTTPResponse {
	feedbacks, err := gs.repository.ReadCheapest()

	if err != nil {
		fmt.Println(err)
		return &response.HTTPResponse{
			Message: "kesalahan saat pengambilan data",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "produk berdasarkan harga termurah berhasil diambil",
		Status:    200,
		IsSuccess: true,
		Data:      feedbacks,
	}
}

// Update implements grocery.GroceryService
func (gs *groceryService) Update(body *request_body.UpdateOrDelete) *response.HTTPResponse {
	if err := gs.repository.Update(body); err != nil {
		return &response.HTTPResponse{
			Message: "kesalahan saat update data",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "produk berhasil diupdate",
		Status:    200,
		IsSuccess: true,
	}
}

func GroceryRepository(infra infrastructures.Infrastructures) grocery.GroceryService {
	return &groceryService{repository: infra.Repositories().Grocery}
}
