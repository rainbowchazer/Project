package dtos

type FeedbackDto struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Purpose string `json:"purpose"`
	Comment string `json:"comment"`
}
