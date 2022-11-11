package service

import (
	"fmt"
	infrastructures "kang-sayur-backend/infrastructure/injector"
	"kang-sayur-backend/model/domain/feedback"
	response "kang-sayur-backend/model/web"
	request_body "kang-sayur-backend/model/web/request_body/feedback"
)

type feedbackService struct {
	repository feedback.FeedbackRepository
}

// Create implements feedback.FeedbackService
func (fs *feedbackService) Create(body *request_body.Create) *response.HTTPResponse {
	inserted, err := fs.repository.Create(body)

	if err != nil {
		return &response.HTTPResponse{
			Message: "kesalahan saat menyimpan data",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "timbal balik berhasil disimpan",
		Status:    200,
		IsSuccess: true,
		Data:      inserted,
	}
}

// Delete implements feedback.FeedbackService
func (fs *feedbackService) Delete(body *request_body.UpdateOrDelete) *response.HTTPResponse {
	if err := fs.repository.Delete(body); err != nil {
		return &response.HTTPResponse{
			Message: "kesalahan saat pengapusan data",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "timbal balik berhasil dihapus",
		Status:    200,
		IsSuccess: true,
	}
}

// Read implements feedback.FeedbackService
func (fs *feedbackService) Read(body *request_body.ReadOnProduct) *response.HTTPResponse {
	feedbacks, err := fs.repository.Read(body)

	if err != nil {
		fmt.Println(err)
		return &response.HTTPResponse{
			Message: "kesalahan saat pengambilan data",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "timbal balik berhasil diambil",
		Status:    200,
		IsSuccess: true,
		Data:      feedbacks,
	}
}

// ReadOne implements feedback.FeedbackService
func (fs *feedbackService) ReadOne(body *request_body.ReadOne) *response.HTTPResponse {
	feedbacks, err := fs.repository.ReadOne(body)

	if err != nil {
		fmt.Println(err)
		return &response.HTTPResponse{
			Message: "kesalahan saat pengambilan data",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "timbal balik berhasil diambil",
		Status:    200,
		IsSuccess: true,
		Data:      feedbacks,
	}
}

// Update implements feedback.FeedbackService
func (fs *feedbackService) Update(body *request_body.UpdateOrDelete) *response.HTTPResponse {
	if err := fs.repository.Update(body); err != nil {
		return &response.HTTPResponse{
			Message: "kesalahan saat update data",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "timbal balik berhasil diupdate",
		Status:    200,
		IsSuccess: true,
	}
}

func FeedbackService(infra *infrastructures.Infrastructures) feedback.FeedbackService {
	return &feedbackService{repository: infra.Repositories().Feedback}
}
