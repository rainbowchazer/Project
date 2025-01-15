package enums

type UserStatus string

const (
	Active       UserStatus = "ACTIVE"
	Banned       UserStatus = "BANNED"
	Confirmation UserStatus = "CONFIRMATION"
)
