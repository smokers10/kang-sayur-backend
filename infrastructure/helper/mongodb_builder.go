package helper

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collections = []string{
	"domicile",
	"price",
	"admin",
	"sub_admin",
	"banner",
	"forgot_password",
	"verification",
	"customer",
	"permission",
	"category",
	"address",
	"invoice",
	"cart",
	"feedback",
	"recipe",
	"invoice_item",
	"recipe_image",
	"recipe_detail",
	"grocery",
	"grocery_images",
}

type mongoBuilder struct {
	database *mongo.Database
}

func MongoBuilder(db *mongo.Database) *mongoBuilder {
	return &mongoBuilder{database: db}
}

func (mb *mongoBuilder) CollectionBuilder() {
	ctx, cancel := InitCtxTimeout()
	defer cancel()

	for _, v := range collections {
		mb.database.CreateCollection(ctx, v)
		mb.collectionIndex(v)
	}
}

func (mb *mongoBuilder) collectionIndex(collection_name string) {
	// for collection that have domicile id
	if collection_name == "price" || collection_name == "customer" {
		model := []mongo.IndexModel{
			{
				Keys:    bson.M{"domicile_id": 1},
				Options: &options.IndexOptions{},
			},
		}

		mb.indexBuilder(collection_name, model)
	}

	// for collection that have customer id
	if collection_name == "address" || collection_name == "invoice" || collection_name == "cart" || collection_name == "feedback" || collection_name == "forgot_password" || collection_name == "verification" {
		model := []mongo.IndexModel{
			{
				Keys:    bson.M{"domicile_id": 1},
				Options: &options.IndexOptions{},
			},
		}

		mb.indexBuilder(collection_name, model)
	}

	// for collection that have recipe id
	if collection_name == "recipe_image" || collection_name == "recipe_detail" {
		model := []mongo.IndexModel{
			{
				Keys:    bson.M{"recipe_id": 1},
				Options: &options.IndexOptions{},
			},
		}

		mb.indexBuilder(collection_name, model)
	}

	// for collection that have product id
	if collection_name == "cart" || collection_name == "feedback" || collection_name == "invoice_item" {
		model := []mongo.IndexModel{
			{
				Keys:    bson.M{"product_id": 1},
				Options: &options.IndexOptions{},
			},
		}

		mb.indexBuilder(collection_name, model)
	}

	// collections that have FK that not related to other collection
	if collection_name == "invoice_item" {
		model := []mongo.IndexModel{
			{
				Keys:    bson.M{"invoice_id": 1},
				Options: &options.IndexOptions{},
			},
		}

		mb.indexBuilder(collection_name, model)
	}

	if collection_name == "grocery_images" {
		model := []mongo.IndexModel{
			{
				Keys:    bson.M{"invoice_id": 1},
				Options: &options.IndexOptions{},
			},
		}

		mb.indexBuilder(collection_name, model)
	}

	if collection_name == "permission" {
		model := []mongo.IndexModel{
			{
				Keys:    bson.M{"sub_admin_id": 1},
				Options: &options.IndexOptions{},
			},
		}

		mb.indexBuilder(collection_name, model)
	}

	if collection_name == "grocery" {
		model := []mongo.IndexModel{
			{
				Keys:    bson.M{"category_id": 1},
				Options: &options.IndexOptions{},
			},
			{
				Keys:    bson.M{"price_id": 1},
				Options: &options.IndexOptions{},
			},
		}

		mb.indexBuilder(collection_name, model)
	}
}

func (mb *mongoBuilder) indexBuilder(collection_name string, model []mongo.IndexModel) {
	ctx, cancel := InitCtxTimeout()
	defer cancel()

	mb.database.Collection(collection_name).Indexes().CreateMany(ctx, model)
}
