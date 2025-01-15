package dtos

type CredentialsDto struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	SubscribedOnFeedback bool   `json:"subscribed_on_feedback"`
}
