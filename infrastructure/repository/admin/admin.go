package admin

import (
	"fmt"
	"kang-sayur-backend/infrastructure/helper"
	"kang-sayur-backend/model/domain/admin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type adminRepository struct {
	collection mongo.Collection
}

// UpdatePassword implements admin.AdminRepository
func (r *adminRepository) UpdatePassword(admin_id string, password string) error {
	ctx, cancel := helper.InitCtxTimeout()
	defer cancel()

	aid, _ := primitive.ObjectIDFromHex(admin_id)
	document := bson.M{
		"$set": bson.M{
			"password": password,
		},
	}

	if _, err := r.collection.UpdateOne(ctx, bson.M{"_id": aid}, document); err != nil {
		return err
	}

	return nil
}

// CheckEmail implements admin.AdminRepository
func (r *adminRepository) CheckEmail(email string) (result *admin.Admin) {
	ctx, cancel := helper.InitCtxTimeout()
	defer cancel()

	if err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&result); err != nil {
		fmt.Println(err)
		return &admin.Admin{}
	}

	return result
}

func AdminRepository(db *mongo.Database) admin.AdminRepository {
	return &adminRepository{
		collection: *db.Collection("admin"),
	}
}
