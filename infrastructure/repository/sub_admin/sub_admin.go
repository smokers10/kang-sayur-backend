package subadmin

import (
	subadmin "kang-sayur-backend/model/domain/sub_admin"
	request_body "kang-sayur-backend/model/web/request_body/sub_admin"

	"go.mongodb.org/mongo-driver/mongo"
)

type subAdminRepository struct {
	collection mongo.Collection
}

// UpdateStatus implements subadmin.SubAdminRepository
func (subAdminRepository) UpdateStatus(data *request_body.UpdateStatus) error {
	panic("unimplemented")
}

// ReadByEmail implements subadmin.SubAdminRepository
func (subAdminRepository) ReadByEmail(email string) (*subadmin.SubAdmin, error) {
	panic("unimplemented")
}

// Create implements subadmin.SubAdminRepository
func (subAdminRepository) Create(data *request_body.Create) error {
	panic("unimplemented")
}

// Delete implements subadmin.SubAdminRepository
func (subAdminRepository) Delete(data *request_body.Delete) error {
	panic("unimplemented")
}

// Read implements subadmin.SubAdminRepository
func (subAdminRepository) Read() ([]subadmin.SubAdmin, error) {
	panic("unimplemented")
}

// Update implements subadmin.SubAdminRepository
func (subAdminRepository) Update(data *request_body.Update) error {
	panic("unimplemented")
}

func SubAdminRepository(db *mongo.Database) subadmin.SubAdminRepository {
	return subAdminRepository{collection: *db.Collection("sub_admin")}
}
