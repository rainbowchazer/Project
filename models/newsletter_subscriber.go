package models

type NewsletterSubscriber struct {
	ID    uint   `grom:"primaryKey;autoIncrement" json:"id"`
	Email string `grom:"not null;unique" json:"email"`
}

func (NewsletterSubscriber) TableName() string {
	return "idel_newsletter_subscribers"
}
