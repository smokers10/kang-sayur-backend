package verification

import (
	"kang-sayur-backend/infrastructure/helper"
	"kang-sayur-backend/model/domain/verification"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type verificationRepository struct {
	collection mongo.Collection
}

// Delete implements verification.VerificationRepository
func (vr *verificationRepository) Delete(customer_id string) error {
	ctx, close := helper.InitCtxTimeout()
	defer close()

	cid, _ := primitive.ObjectIDFromHex(customer_id)

	if err := vr.collection.FindOneAndDelete(ctx, bson.M{"customer_id": cid}).Err(); err != nil {
		return err
	}

	return nil
}

// ReadOne implements verification.VerificationRepository
func (vr *verificationRepository) ReadOne(token string, customer_id string) (result *verification.Verification) {
	ctx, close := helper.InitCtxTimeout()
	defer close()

	cid, _ := primitive.ObjectIDFromHex(customer_id)

	vr.collection.FindOne(ctx, bson.M{"token": token, "customer_id": cid}).Decode(&result)

	return result
}

// Upsert implements verification.VerificationRepository
func (vr *verificationRepository) Upsert(token string, customer_id string, verification_code string) error {
	ctx, close := helper.InitCtxTimeout()
	defer close()

	cid, _ := primitive.ObjectIDFromHex(customer_id)
	document := bson.M{
		"$set": bson.M{
			"token":             token,
			"verification_code": verification_code,
		},
	}

	if _, err := vr.collection.UpdateOne(ctx, bson.M{"customer_id": cid}, document, options.Update().SetUpsert(true)); err != nil {
		return err
	}

	return nil
}

func VerificationRepository(db *mongo.Database) verification.VerificationRepository {
	return &verificationRepository{
		collection: *db.Collection("verification"),
	}
}
