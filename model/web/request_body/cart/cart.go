package cart

// BasicAction : for creating and deleting
type BasicAction struct {
	CustomerID string `json:"customer_id" bson:"customer_id"`
	ProductID  string `json:"product_id" bson:"product_id"`
	Quantity   int    `json:"quantity" bson:"quantity"`
}

type UpdateQuantity struct {
	ID       string `json:"id" bson:"_id"`
	Quantity int    `json:"quantity" bson:"quantity"`
}
