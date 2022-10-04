package grocery

type Create struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	UnitType     string `json:"unit_type"`
	PointOfSales string `json:"point_of_sales"`
	CategoryID   string `json:"category_id"`
	PriceID      string `json:"price_id"`
}

type UpdateOrDelete struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	UnitType     string `json:"unit_type"`
	PointOfSales string `json:"point_of_sales"`
	CategoryID   string `json:"category_id"`
	PriceID      string `json:"price_id"`
}

type ByCategory struct {
	CategoryID string `json:"category_id"`
}

type ByKeyword struct {
	Keyword string `json:"keyword"`
}
