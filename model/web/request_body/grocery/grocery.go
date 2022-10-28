package grocery

type Create struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	UnitType     string `json:"unit_type"`
	PointOfSales string `json:"point_of_sales"`
	CategoryID   string `json:"category_id"`
	Price        string `json:"price"`
}

type UpdateOrDelete struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	UnitType     string `json:"unit_type"`
	PointOfSales string `json:"point_of_sales"`
	CategoryID   string `json:"category_id"`
	Price        string `json:"price"`
}

type ByCategory struct {
	CategoryID string `json:"category_id"`
}

type ByKeyword struct {
	Keyword string `json:"keyword"`
}

type SetOrUpdatePrice struct {
	ID         string `json:"id"`
	Price      int    `json:"price"`
	DomicileID string `json:"domicile_id"`
}

type CreateOrDeleteGroceryImage struct {
	ID        string `json:"id"`
	Source    string `json:"source"`
	GroceryID string `json:"grocery_id"`
}
