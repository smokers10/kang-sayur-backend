package feedback

type Create struct {
	CustomerID string `json:"customer_id" bson:"customer_id"`
	ProductID  string `json:"product_id" bson:"product_id"`
	Body       string `json:"body" bson:"body"`
	Rating     string `json:"rating" bson:"rating"`
}

type UpdateOrDelete struct {
	ID         string `json:"id" bson:"_id"`
	CustomerID string `json:"customer_id" bson:"customer_id"`
	ProductID  string `json:"product_id" bson:"product_id"`
	Body       string `json:"body" bson:"body"`
	Rating     string `json:"rating" bson:"rating"`
}

type ReadOne struct {
	ID string `json:"id" bson:"_id"`
}

type ReadOnProduct struct {
	ID string `json:"id" bson:"_id"`
}
