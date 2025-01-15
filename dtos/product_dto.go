package dtos

import (
	"idel/models/seo"
	"mime/multipart"
)

type ProductDto struct {
	ID              uint    `json:"id"`
	Seo             seo.Seo `json:"seo"`
	ShortName       string  `json:"short_name"`
	FullName        string  `json:"full_name"`
	Description     string  `json:"description"`
	Usage           string  `json:"usage"`
	Precautions     string  `json:"precautions"`
	GuaranteePeriod string  `json:"guarantee_period"`
	Composition     string  `json:"composition"`
	Purpose         string  `json:"purpose"`
	VendorCode      string  `json:"vendor_code"`
	LinkSweetlife   string  `json:"link_sweetlife"`
	LinkWildberries string  `json:"link_wildberries"`
	LinkOzon        string  `json:"link_ozon"`

	CategoryMainID             *uint                   `json:"category_main_id"`
	CategoryWhatID             *uint                   `json:"category_what_id"`
	CategoryForID              *uint                   `json:"category_for_id"`
	DockInstructionFilePaths   []string                `json:"dock_instruction_file_paths"`
	DockCertificationFilePaths []string                `json:"dock_certification_file_paths"`
	ImageFilePaths             []string                `json:"image_file_paths"`
	DockCertificationFiles     []*multipart.FileHeader `json:"dock_certification_files"`
	DockInstructionFiles       []*multipart.FileHeader `json:"dock_instruction_files"`
	ImageFiles                 []*multipart.FileHeader `json:"image_files"`
}
