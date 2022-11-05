package permission

import (
	"kang-sayur-backend/model/domain/permission"
	request_body "kang-sayur-backend/model/web/request_body/sub_admin"

	"go.mongodb.org/mongo-driver/mongo"
)

type permissionRepository struct {
	collection mongo.Collection
}

// Upsert implements permission.PermissionRepository
func (*permissionRepository) Upsert(data *request_body.SetPermission) error {
	panic("unimplemented")
}

// ReadOne implements permission.PermissionRepository
func (*permissionRepository) ReadOne(sub_admin_id string) (*permission.Permission, error) {
	panic("unimplemented")
}

func PermissionRepository(db *mongo.Database) permission.PermissionRepository {
	return &permissionRepository{collection: *db.Collection("permission")}
}
