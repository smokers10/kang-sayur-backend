package customer

type Register struct {
	Name                 string `json:"name" bson:"name"`
	Phone                string `json:"phone" bson:"phone"`
	Email                string `json:"email" bson:"email"`
	Password             string `json:"password" bson:"password"`
	PasswordConfirmation string `json:"password_confirmation" bson:"password_confirmation"`
	DomicileID           string `json:"domicile_id" bson:"domicile_id"`
}

type UpdateProfile struct {
	CustomerID string `json:"customer_id" bson:"customer_id"`
	Name       string `json:"name" bson:"name"`
	Phone      string `json:"phone" bson:"phone"`
	DomicileID string `json:"domicile_id" bson:"domicile_id"`
}

type Login struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type ForgotPassword struct {
	Email string `json:"email" bson:"email"`
}

type ResetPassword struct {
	Token                string `json:"token" bson:"token"`
	Code                 string `json:"code" bson:"code"`
	Password             string `json:"password" bson:"password"`
	PasswordConfirmation string `json:"password_confirmation" bson:"password_confirmation"`
}

type ValidateAccount struct {
	CustomerID string `json:"customer_id" bson:"customer_id"`
	Code       string `json:"code" bson:"code"`
}

type ViewProfile struct {
	Id string `json:"id"`
}
