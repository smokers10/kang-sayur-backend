package middleware

type MiddlewareResponse struct {
	Message string `json:"message,omitempty"`
	Status  int    `json:"status,omitempty"`
	Is_pass bool   `json:"is_password,omitempty"`
	Reason  string `json:"reason,omitempty"`
	Claim   struct {
		Id    string `json:"id,omitempty"`
		Email string `json:"email,omitempty"`
	}
	Permission interface{}
}

type Contract interface {
	Process(token string) *MiddlewareResponse
}
