package dtos

type RecaptchaResponseDto struct {
	Success     bool     `json:"success"`               // Статус успешности проверки
	ChallengeTs string   `json:"challenge_ts"`          // Время проверки
	Action      string   `json:"action,omitempty"`      // Действие (если есть)
	Hostname    string   `json:"hostname"`              // Хост, на котором была выполнена проверка
	Score       float64  `json:"score,omitempty"`       // Оценка риска (для reCAPTCHA v3)
	ErrorCodes  []string `json:"error-codes,omitempty"` // Ошибки, если есть
}
