package dtos

import (
	"idel/models/seo"
	"mime/multipart"
	"time"
)

type NewsDto struct {
	Id               uint                  `json:"id"`
	Seo              seo.Seo               `json:"seo"`
	NewsletterEnable bool                  `json:"newsletter_enable"`
	Title            string                `json:"title"`
	Date             time.Time             `json:"date"`
	Text             string                `json:"text"`
	ImagePath        string                `json:"image_path"`
	Image            *multipart.FileHeader `json:"image"`
}
