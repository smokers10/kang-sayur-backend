package forgotpassword

import (
	"kang-sayur-backend/infrastructure/helper"
	forgotpassword "kang-sayur-backend/model/domain/forgot_password"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type forgotPassword struct {
	collection mongo.Collection
}

// Delete implements forgotpassword.ForgotPasswordRepository
func (fp *forgotPassword) Delete(token string) error {
	ctx, close := helper.InitCtxTimeout()
	defer close()

	if err := fp.collection.FindOneAndDelete(ctx, bson.M{"token": token}).Err(); err != nil {
		return err
	}

	return nil
}

// ReadOne implements forgotpassword.ForgotPasswordRepository
func (fp *forgotPassword) ReadOne(token string) (result *forgotpassword.ForgotPassword) {
	ctx, close := helper.InitCtxTimeout()
	defer close()

	fp.collection.FindOne(ctx, bson.M{"token": token}).Decode(&result)

	return result
}

// Upsert implements forgotpassword.ForgotPasswordRepository
func (fp *forgotPassword) Upsert(token string, user_id string, code string) error {
	ctx, close := helper.InitCtxTimeout()
	defer close()

	cid, _ := primitive.ObjectIDFromHex(user_id)

	document := bson.M{
		"$set": bson.M{
			"token":   token,
			"user_id": cid,
			"code":    code,
		},
	}

	if _, err := fp.collection.UpdateOne(ctx, bson.M{"user_id": user_id}, document); err != nil {
		return err
	}

	return nil
}

func ForgotPassword(db *mongo.Database) forgotpassword.ForgotPasswordRepository {
	return &forgotPassword{
		collection: *db.Collection("forgot_password"),
	}
}
