package permission

import (
	"kang-sayur-backend/model/domain/permission"

	"go.mongodb.org/mongo-driver/mongo"
)

type permissionRepository struct {
	collection mongo.Collection
}

// ReadOne implements permission.PermissionRepository
func (*permissionRepository) ReadOne(sub_admin_id string) (*permission.Permission, error) {
	panic("unimplemented")
}

// Upsert implements permission.PermissionRepository
func (*permissionRepository) Upsert(data *permission.Permission) error {
	panic("unimplemented")
}

func PermissionRepository(db *mongo.Database) permission.PermissionRepository {
	return &permissionRepository{collection: *db.Collection("permission")}
}
