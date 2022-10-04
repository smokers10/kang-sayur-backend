package category

type CreateCategory struct {
	Name        string `json:"name"`
	Ilustration string `json:"Ilustration"`
}

type UpdateOrDeleteCategory struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Ilustration string `json:"Ilustration"`
}
