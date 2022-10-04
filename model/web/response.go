package web

type HTTPResponse struct {
	Message   string      `json:"message"`
	Status    int         `json:"status"`
	IsSuccess bool        `json:"is_success"`
	Data      interface{} `json:"data"`
	Token     string      `json:"token,omitempty"`
}
