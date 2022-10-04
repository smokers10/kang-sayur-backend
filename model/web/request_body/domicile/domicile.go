package domicile

type Create struct {
	Name string `json:"name"`
}

type UpdateOrDelete struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
