package forgotpassword

type Request struct {
	Email string `json:"email"`
}

type ResetPassword struct {
	Token  string `json:"token"`
	UserID string `json:"user_id"`
	Code   string `json:"code"`
}
