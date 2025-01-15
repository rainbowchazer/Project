package models

import (
	"idel/models/seo"

	"github.com/lib/pq"
)

type Product struct {
	ID              uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Seo             seo.Seo `gorm:"embedded" json:"seo"`
	ShortName       string  `gorm:"column:short_name;not null" json:"short_name"`
	FullName        string  `gorm:"column:full_name;not null" json:"full_name"`
	Description     string  `json:"description"`
	Usage           string  `json:"usage"`
	Precautions     string  `json:"precautions"`
	GuaranteePeriod string  `gorm:"not null;column:guarantee_period" json:"guarantee_period"`
	Composition     string  `json:"composition"`
	Purpose         string  `json:"purpose"`
	VendorCode      string  `gorm:"column:vendor_code" json:"vendor_code"`
	LinkSweetlife   string  `gorm:"column:link_sweetlife" json:"link_sweetlife"`
	LinkWildberries string  `gorm:"column:link_wildberries" json:"link_wildberries"`
	LinkOzon        string  `gorm:"column:link_ozon" json:"link_ozon"`

	CategoryMainID             *uint          `gorm:"column:category_main_id" json:"category_main_id"`
	CategoryMain               *Category      `gorm:"column:category_main" json:"category_main"`
	CategoryWhatID             *uint          `gorm:"category_what_id" json:"category_what_id"`
	CategoryWhat               *Category      `gorm:"column:category_what" json:"category_what"`
	CategoryForID              *uint          `gorm:"category_for_id" json:"category_for_id"`
	CategoryFor                *Category      `gorm:"column:category_for" json:"category_for"`
	DockInstructionFilePaths   pq.StringArray `gorm:"type:text[]" column:"dock_instruction_file_names" json:"dock_instruction_file_paths"`
	DockCertificationFilePaths pq.StringArray `gorm:"type:text[]" column:"dock_certification_file_names" json:"dock_certification_file_paths"`
	ImageFilePaths             pq.StringArray `gorm:"type:text[]" column:"image_file_names" json:"image_file_paths"`
}

func (Product) TableName() string {
	return "idel_products"
}

func (p *Product) GetFullName() string {
	return p.FullName
}

func (p *Product) GetSeoIdSource() string {
	return p.GetFullName()
}
