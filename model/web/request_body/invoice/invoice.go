package invoice

type Checkout struct {
	PaymentMethod string `json:"payment_method"`
	PaymentToken  string `json:"payment_token"`
	CustomerID    string `json:"customer_id"`
	Items         []Item `json:"items"`
}

type Item struct {
	ProductID   string `json:"product_id"`
	ProductType string `json:"product_type"`
	Quantities  string `json:"quantities"`
}

type UpdateStatus struct {
	InvoiceID string `json:"invoice_id"`
	Status    string `json:"status"`
}

type ReadOne struct {
	InvoiceID string `json:"invoice_id"`
	Status    string `json:"status"`
}

type Pay struct {
	PaymentToken string `json:"PaymentToken"`
}
