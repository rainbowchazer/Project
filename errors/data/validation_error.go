package data

import (
	"fmt"
)

// ValidationViolationDto - структура для хранения информации о нарушении валидации
type ValidationViolationDto struct {
	Field       string
	Message     string
	Value       interface{}
	Constraints map[string]interface{}
}

// NewValidationViolationDto - создает новый объект ValidationViolationDto
func NewValidationViolationDto(field string, message string) *ValidationViolationDto {
	return &ValidationViolationDto{
		Field:   field,
		Message: message,
	}
}

// NewValidationViolationDtoWithValue - создает новый объект ValidationViolationDto с значением
func NewValidationViolationDtoWithValue(field string, value interface{}, message string) *ValidationViolationDto {
	return &ValidationViolationDto{
		Field:   field,
		Value:   value,
		Message: message,
	}
}

// NewValidationViolationDtoWithConstraints - создает новый объект ValidationViolationDto с значением и ограничениями
func NewValidationViolationDtoWithConstraints(field string, value interface{}, message string, constraints map[string]interface{}) *ValidationViolationDto {
	return &ValidationViolationDto{
		Field:       field,
		Value:       value,
		Message:     message,
		Constraints: constraints,
	}
}

// BuildRequiredFieldNotPresent - статический метод для создания нарушения валидации с сообщением о поле, которое не может быть пустым
func BuildRequiredFieldNotPresent(field string, messageStart string) *ValidationViolationDto {
	message := fmt.Sprintf("%s must not be null", messageStart)
	return NewValidationViolationDto(field, message)
}

// ValidationException - структура ошибки с дополнительными деталями (сообщения и нарушения валидации)
type ValidationError struct {
	Message string
	Details map[string]interface{}
}

// NewValidationException - создает новый объект ValidationException
func NewValidationException(message string, details map[string]interface{}) *ValidationError {
	return &ValidationError{
		Message: message,
		Details: details,
	}
}

// ValidationExceptionBuilder - строит ValidationException с использованием шаблона Builder
type ValidationErrorBuilder struct {
	message    string
	violations []*ValidationViolationDto
	details    map[string]interface{}
}

// NewValidationExceptionBuilder - создает новый объект ValidationExceptionBuilder
func NewValidationExceptionBuilder() *ValidationErrorBuilder {
	return &ValidationErrorBuilder{
		details: make(map[string]interface{}),
	}
}

// AddViolation - добавляет нарушение в список
func (b *ValidationErrorBuilder) AddViolation(dto *ValidationViolationDto) *ValidationErrorBuilder {
	b.violations = append(b.violations, dto)
	return b
}

// AddDetails - добавляет дополнительную информацию в детали ошибки
func (b *ValidationErrorBuilder) AddDetails(key string, value interface{}) *ValidationErrorBuilder {
	b.details[key] = value
	return b
}

// SetMessage - устанавливает сообщение для ошибки
func (b *ValidationErrorBuilder) SetMessage(message string) *ValidationErrorBuilder {
	b.message = message
	return b
}

// HasViolations - проверяет, есть ли нарушения валидации
func (b *ValidationErrorBuilder) HasViolations() bool {
	return len(b.violations) > 0
}

// Build - создает новый объект ValidationException с применением собранных данных
func (b *ValidationErrorBuilder) Build() *ValidationError {
	if b.HasViolations() {
		b.details["violations"] = b.violations
	}
	return &ValidationError{
		Message: b.message,
		Details: b.details,
	}
}
