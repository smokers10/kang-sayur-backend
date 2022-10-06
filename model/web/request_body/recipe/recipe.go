package recipe

type Create struct {
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	UnitType     string  `json:"unit_type"`
	PointOfSales string  `json:"point_of_sales"`
	HowToCook    string  `json:"how_to_cook"`
	DeleteStatus string  `json:"delete_status"`
	CategoryID   string  `json:"category_id"`
	PriceID      string  `json:"price_id"`
	Images       []Image `json:"images"`
}

type Update struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	UnitType     string `json:"unit_type"`
	PointOfSales string `json:"point_of_sales"`
	HowToCook    string `json:"how_to_cook"`
	DeleteStatus string `json:"delete_status"`
	CategoryID   string `json:"category_id"`
	PriceID      string `json:"price_id"`
}

type RecipeDetail struct {
	ID string `json:"id"`
}

type Delete struct {
	ID string `json:"id"`
}

type Image struct {
	Base64 string `json:"base64"`
	Format string `json:"format"`
}

type DeleteImage struct {
	ID string `json:"id"`
}

type UpdateOrderImage struct {
	ID    string `json:"id"`
	Order int    `json:"order"`
}

type Searching struct {
	Keyword string `json:"keyword"`
}

type AddDetail struct {
	Name     string `json:"name"`
	UnitType string `json:"unit_type"`
	RecipeID string `json:"recipe_id"`
}

type UpdateOrDeleteDetail struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	UnitType string `json:"unit_type"`
	RecipeID string `json:"recipe_id"`
}
