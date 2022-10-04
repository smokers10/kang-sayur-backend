package address

type CreateAddress struct {
	Name    string `json:"name" bson:"name"`
	Address string `json:"address" bson:"address"`
	Status  string `json:"status" bson:"status"`
}

type UpdateAddress struct {
	ID         string `json:"id" bson:"_id"`
	Name       string `json:"name" bson:"name"`
	Address    string `json:"address" bson:"address"`
	Status     string `json:"status" bson:"status"`
	CustomerID string `json:"customer_id" bson:"customer_id"`
}

type DeleteOrReadAddress struct {
	ID         string `json:"id" bson:"_id"`
	CustomerID string `json:"customer_id" bson:"customer_id"`
}

type ReadOne struct {
	ID string `json:"id" bson:"_id"`
}
