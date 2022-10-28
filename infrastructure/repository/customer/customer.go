package customer

import (
	"kang-sayur-backend/infrastructure/helper"
	"kang-sayur-backend/model/domain/customer"
	request_body "kang-sayur-backend/model/web/request_body/customer"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type customerRepository struct {
	collection *mongo.Collection
}

// Create implements customer.CustomerRepository
func (cr *customerRepository) Create(data *request_body.Register) (result *customer.Customer, err error) {
	ctx, close := helper.InitCtxTimeout()
	defer close()

	did, _ := primitive.ObjectIDFromHex(data.DomicileID)

	document := bson.M{
		"name":        data.Name,
		"phone":       data.Phone,
		"email":       data.Email,
		"password":    data.Password,
		"domicile_id": did,
	}

	r, err := cr.collection.InsertOne(ctx, document)
	if err != nil {
		return nil, err
	}

	res := customer.Customer{
		ID:                 r.InsertedID.(primitive.ObjectID).Hex(),
		Name:               data.Name,
		Phone:              data.Phone,
		Email:              data.Email,
		VerificationStatus: "no verified",
		DomicileID:         did.Hex(),
	}

	return &res, nil
}

// Read implements customer.CustomerRepository
func (cr *customerRepository) Read() (result []customer.Customer) {
	ctx, close := helper.InitCtxTimeout()
	defer close()

	cur, _ := cr.collection.Find(ctx, bson.M{})

	cur.All(ctx, result)

	return result
}

// ReadByEmail implements customer.CustomerRepository
func (cr *customerRepository) ReadByEmail(email string) (result *customer.Customer) {
	ctx, close := helper.InitCtxTimeout()
	defer close()

	cr.collection.FindOne(ctx, bson.M{"email": email}).Decode(&result)

	return result
}

// ReadByID implements customer.CustomerRepository
func (cr *customerRepository) ReadByID(id string) (result *customer.Customer) {
	ctx, close := helper.InitCtxTimeout()
	defer close()

	_id, _ := primitive.ObjectIDFromHex(id)

	cr.collection.FindOne(ctx, bson.M{"_id": _id}).Decode(&result)

	panic("unimplemented")
}

// UpdatePassword implements customer.CustomerRepository
func (cr *customerRepository) UpdatePassword(password string, customer_id string) error {
	ctx, close := helper.InitCtxTimeout()
	defer close()

	_id, _ := primitive.ObjectIDFromHex(customer_id)

	if err := cr.collection.FindOneAndUpdate(ctx, bson.M{"_id": _id}, bson.M{"password": password}).Err(); err != nil {
		return err
	}

	return nil
}

// UpdateProfile implements customer.CustomerRepository
func (cr *customerRepository) UpdateProfile(data *request_body.UpdateProfile) error {
	ctx, close := helper.InitCtxTimeout()
	defer close()

	_id, _ := primitive.ObjectIDFromHex(data.CustomerID)

	document := bson.M{
		"$set": bson.M{
			"name":        data.Name,
			"phone":       data.Phone,
			"domicile_id": data.DomicileID,
		},
	}

	if err := cr.collection.FindOneAndUpdate(ctx, bson.M{"_id": _id}, document).Err(); err != nil {
		return err
	}

	return nil
}

// VerifyVerification implements customer.CustomerRepository
func (cr *customerRepository) VerifyVerification(customer_id string) error {
	ctx, close := helper.InitCtxTimeout()
	defer close()

	_id, _ := primitive.ObjectIDFromHex(customer_id)

	if err := cr.collection.FindOneAndUpdate(ctx, bson.M{"_id": _id}, bson.M{"verification_status": "verified"}).Err(); err != nil {
		return err
	}

	return nil
}

func CustomerRepository(db *mongo.Database) customer.CustomerRepository {
	return &customerRepository{
		collection: db.Collection("customer"),
	}
}
