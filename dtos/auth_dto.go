package dtos

import (
	"idel/enums"
)

type AuthDto struct {
	Email                string         `json:"email"`
	Token                string         `json:"token"`
	Role                 enums.UserRole `json:"role"`
	SubscribedOnFeedback bool           `json:"subscribed_on_feedback"`
}
