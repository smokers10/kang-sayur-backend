package cart

// BasicAction : for creating and deleting
type BasicAction struct {
	CustomerID string `json:"customer_id"`
	ProductID  string `json:"product_id"`
	Quantity   int    `json:"quantity"`
}

type UpdateQuantity struct {
	ID       string `json:"id" bson:"_id"`
	Quantity int    `json:"quantity"`
}

type ReadCart struct {
	CustomerID string `json:"customer_id"`
}
