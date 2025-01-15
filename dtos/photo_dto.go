package dtos

import (
	"mime/multipart"
)

// PhotoDto представляет данные фотографии для передачи между слоями
type PhotoDto struct {
	ID          uint64                `json:"id,omitempty"` // Идентификатор фотографии (опционально)
	Description string                `json:"description"`  // Описание фотографии
	ImagePath   string                `json:"image_path"`   // Путь к изображению
	Image       *multipart.FileHeader `json:"image"`        // Загружаемый файл (не сериализуется в JSON)
}
