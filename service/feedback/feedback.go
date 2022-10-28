package service

import (
	infrastructures "kang-sayur-backend/infrastructure/injector"
	"kang-sayur-backend/model/domain/feedback"
	response "kang-sayur-backend/model/web"
	request_body "kang-sayur-backend/model/web/request_body/feedback"
)

type feedbackService struct {
	feedback feedback.FeedbackRepository
}

// Create implements feedback.FeedbackService
func (*feedbackService) Create(body *request_body.Create) *response.HTTPResponse {
	panic("unimplemented")
}

// Delete implements feedback.FeedbackService
func (*feedbackService) Delete(body *request_body.UpdateOrDelete) *response.HTTPResponse {
	panic("unimplemented")
}

// Read implements feedback.FeedbackService
func (*feedbackService) Read(body *request_body.ReadOnProduct) *response.HTTPResponse {
	panic("unimplemented")
}

// ReadOne implements feedback.FeedbackService
func (*feedbackService) ReadOne(body *request_body.ReadOne) *response.HTTPResponse {
	panic("unimplemented")
}

// Update implements feedback.FeedbackService
func (*feedbackService) Update(body *request_body.UpdateOrDelete) *response.HTTPResponse {
	panic("unimplemented")
}

func FeedbackService(infra *infrastructures.Infrastructures) feedback.FeedbackService {
	return &feedbackService{feedback: infra.Repositories().Feedback}
}
