package helper

import "go.mongodb.org/mongo-driver/mongo"

type MongoBuilder struct{}

func (mb *MongoBuilder) CollectionBuilder(db *mongo.Database) {
	ctx, cancel := InitCtxTimeout()
	defer cancel()

	collections := []string{
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

	for _, v := range collections {
		db.CreateCollection(ctx, v)
	}
}

func (mb *MongoBuilder) CollectionIndex(db *mongo.Database, collection_name string) {

}
