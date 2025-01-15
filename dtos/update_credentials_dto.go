package dtos

type UpdateCredentialsDto struct {
	OldEmail             string `json:"old_email"`
	NewEmail             string `json:"new_email"`
	Password             string `json:"password"`
	SubscribedOnFeedback bool   `json:"subscribed_on_feedback"`
}
