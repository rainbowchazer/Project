package event

type CategoryDeleted struct {
	CategoryID uint
}

func NewCategoryDeleted(id uint) *CategoryDeleted {
	return &CategoryDeleted{
		CategoryID: id,
	}
}
