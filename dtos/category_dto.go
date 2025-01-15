package dtos

import (
	"idel/models/seo"
	"mime/multipart"
)

type CategoryDto struct {
	ID           uint                  `json:"id"`
	Seo          seo.Seo               `json:"seo"`
	ParentID     *uint                 `json:"parent_id"`
	ShowInSlider bool                  `json:"show_in_slider"`
	Name         string                `json:"name"`
	Description  string                `json:"description"`
	ImagePath    string                `json:"image_path"`
	Image        *multipart.FileHeader `json:"image"`
}
