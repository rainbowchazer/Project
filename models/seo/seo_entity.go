package seo

// ISeoEntity описывает сущности, содержащие SEO-поля.
type ISeoEntity interface {
	GetSeoIdSource() string // Получить идентификатор SEO (например, имя или уникальный URL)
	GetId() any             // Получить идентификатор сущности (тип может быть разным, поэтому используется any)
	GetSeo() Seo            // Получить структуру SEO
	SetSeo(seo Seo)         // Установить структуру SEO
}
