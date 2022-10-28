package category

import (
	"kang-sayur-backend/model/domain/category"
	request_body "kang-sayur-backend/model/web/request_body/category"

	"go.mongodb.org/mongo-driver/mongo"
)

type categoryRepository struct {
	collection mongo.Collection
}

// Create implements category.CategoryRepository
func (*categoryRepository) Create(body request_body.CreateCategory) (*category.Category, error) {
	panic("unimplemented")
}

// Delete implements category.CategoryRepository
func (*categoryRepository) Delete(body request_body.UpdateOrDeleteCategory) error {
	panic("unimplemented")
}

// Read implements category.CategoryRepository
func (*categoryRepository) Read() ([]category.Category, error) {
	panic("unimplemented")
}

// Update implements category.CategoryRepository
func (*categoryRepository) Update(body request_body.UpdateOrDeleteCategory) error {
	panic("unimplemented")
}

func CategoryRepository(db *mongo.Database) category.CategoryRepository {
	return &categoryRepository{collection: *db.Collection("category")}
}
