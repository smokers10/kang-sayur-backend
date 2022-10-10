package feedback

import (
	response "kang-sayur-backend/model/web"
	request_body "kang-sayur-backend/model/web/request_body/feedback"
)

type Feedback struct {
	ID         string `json:"id" bson:"_id"`
	CustomerID string `json:"customer_id" bson:"customer_id"`
	ProductID  string `json:"product_id" bson:"product_id"`
	Body       string `json:"body" bson:"body"`
	Rating     string `json:"rating" bson:"rating"`
}

type FeedbackService interface {
	Create(body *request_body.Create) *response.HTTPResponse

	Read(body *request_body.ReadOnProduct) *response.HTTPResponse

	ReadOne(body *request_body.ReadOne) *response.HTTPResponse

	Update(body *request_body.UpdateOrDelete) *response.HTTPResponse

	Delete(body *request_body.UpdateOrDelete) *response.HTTPResponse
}

type FeedbackRepository interface {
	Create(data *request_body.Create) error

	Read(data *request_body.ReadOnProduct) ([]Feedback, error)

	ReadOne(data *request_body.ReadOne) (*Feedback, error)

	Update(data *request_body.UpdateOrDelete) error

	Delete(data *request_body.UpdateOrDelete) error
}
