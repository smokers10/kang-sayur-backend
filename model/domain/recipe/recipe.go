package domain

type Recipe struct {
	ID           string `json:"id" bson:"_id"`
	Name         string `json:"name" bson:"name"`
	Description  string `json:"description" bson:"description"`
	UnitType     string `json:"unit_type" bson:"unit_type"`
	PointOfSales string `json:"point_of_sales" bson:"point_of_sales"`
	HowToCook    string `json:"how_to_cook" bson:"how_to_cook"`
	DeleteStatus string `json:"delete_status" bson:"delete_status"`
	CategoryID   string `json:"category_id" bson:"category_id"`
	PriceID      string `json:"price_id" bson:"price_id"`
}
