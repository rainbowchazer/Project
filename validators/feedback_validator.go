package validators

import (
	"errors"
	"idel/dtos"
)

type FeedbackValidator struct {
	Email string `json:"email"`
}

func NewFeedbackValidator() *FeedbackValidator {
	return &FeedbackValidator{}
}

func (v *FeedbackValidator) Validate(feedbackDto dtos.FeedbackDto) error {
	if feedbackDto.Email == "" {
		return errors.New("email cannot be null or empty")
	}
	return nil
}
