package verification

type RequestVerification struct {
	CustomerID string `json:"customer_id"`
}

type Verify struct {
	Token            string `json:"token"`
	CustomerID       string `json:"customer_id"`
	VerificationCode string `json:"Verification_code"`
}
