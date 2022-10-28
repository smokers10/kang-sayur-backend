package service

import (
	infrastructures "kang-sayur-backend/infrastructure/injector"
	"kang-sayur-backend/model/domain/grocery"
	response "kang-sayur-backend/model/web"
	request_body "kang-sayur-backend/model/web/request_body/grocery"
)

type groceryService struct {
	repository grocery.GroceryRepository
}

// Create implements grocery.GroceryService
func (*groceryService) Create(body *request_body.Create) *response.HTTPResponse {
	panic("unimplemented")
}

// Delete implements grocery.GroceryService
func (*groceryService) Delete(body *request_body.UpdateOrDelete) *response.HTTPResponse {
	panic("unimplemented")
}

// Read implements grocery.GroceryService
func (*groceryService) Read(body *request_body.Create) *response.HTTPResponse {
	panic("unimplemented")
}

// ReadBestSeller implements grocery.GroceryService
func (*groceryService) ReadBestSeller() *response.HTTPResponse {
	panic("unimplemented")
}

// ReadByCategory implements grocery.GroceryService
func (*groceryService) ReadByCategory(body *request_body.ByCategory) *response.HTTPResponse {
	panic("unimplemented")
}

// ReadByKeyword implements grocery.GroceryService
func (*groceryService) ReadByKeyword(body *request_body.ByKeyword) *response.HTTPResponse {
	panic("unimplemented")
}

// ReadCheapest implements grocery.GroceryService
func (*groceryService) ReadCheapest() *response.HTTPResponse {
	panic("unimplemented")
}

// SetPrice implements grocery.GroceryService
func (*groceryService) SetPrice(body *request_body.SetOrUpdatePrice) *response.HTTPResponse {
	panic("unimplemented")
}

// Update implements grocery.GroceryService
func (*groceryService) Update(body *request_body.UpdateOrDelete) *response.HTTPResponse {
	panic("unimplemented")
}

// UpdatePrice implements grocery.GroceryService
func (*groceryService) UpdatePrice(body *request_body.SetOrUpdatePrice) *response.HTTPResponse {
	panic("unimplemented")
}

func GroceryRepository(infra infrastructures.Infrastructures) grocery.GroceryService {
	return &groceryService{repository: infra.Repositories().Grocery}
}
