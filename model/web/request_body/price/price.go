package price

type Data struct {
	ID         string `json:"id"`
	Price      int    `json:"price"`
	DomicileID string `json:"domicile_id"`
}

type ByDomicile struct {
	DomicileID string `json:"domicile_id"`
}
