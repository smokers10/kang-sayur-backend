package feedback

import (
	"kang-sayur-backend/model/domain/feedback"
	request_body "kang-sayur-backend/model/web/request_body/feedback"

	"go.mongodb.org/mongo-driver/mongo"
)

type feedbackRepository struct {
	collection mongo.Collection
}

// Create implements feedback.FeedbackRepository
func (*feedbackRepository) Create(data *request_body.Create) error {
	panic("unimplemented")
}

// Delete implements feedback.FeedbackRepository
func (*feedbackRepository) Delete(data *request_body.UpdateOrDelete) error {
	panic("unimplemented")
}

// Read implements feedback.FeedbackRepository
func (*feedbackRepository) Read(data *request_body.ReadOnProduct) ([]feedback.Feedback, error) {
	panic("unimplemented")
}

// ReadOne implements feedback.FeedbackRepository
func (*feedbackRepository) ReadOne(data *request_body.ReadOne) (*feedback.Feedback, error) {
	panic("unimplemented")
}

// Update implements feedback.FeedbackRepository
func (*feedbackRepository) Update(data *request_body.UpdateOrDelete) error {
	panic("unimplemented")
}

func FeedbackRepository(db *mongo.Database) feedback.FeedbackRepository {
	return &feedbackRepository{collection: *db.Collection("feedback")}
}
