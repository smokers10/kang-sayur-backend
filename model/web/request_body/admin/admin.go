package requestbody

type LoginRequest struct {
	Email string `json:"email"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
