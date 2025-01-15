package seo

type Seo struct {
	ID       string       `gorm:"column:seo_id;unique;not null"`
	URL      string       `gorm:"column:seo_url;unique;not null"`
	MetaTags []SeoMetaTag `gorm:"embedded;embeddedPrefix:meta_"`
}
