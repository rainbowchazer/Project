package data

import (
	"idel/dtos"
	"strings"
)

type RecaptchaError struct {
	Message string
}

func NewRecaptchaError(response dtos.RecaptchaResponseDto) *RecaptchaError {
	return &RecaptchaError{
		Message: buildErrorMessage(response),
	}
}

func (e *RecaptchaError) Error() string {
	return e.Message
}

func buildErrorMessage(response dtos.RecaptchaResponseDto) string {
	var outputMessage strings.Builder
	outputMessage.WriteString("reCAPTCHA errors occurred:")

	if response.ErrorCodes != nil {
		for _, errorCode := range response.ErrorCodes {
			switch errorCode {
			case "missing-input-secret":
				outputMessage.WriteString(" The secret parameter is missing;")
			case "invalid-input-secret":
				outputMessage.WriteString(" The secret parameter is invalid or malformed;")
			case "missing-input-response":
				outputMessage.WriteString(" The response parameter is missing;")
			case "invalid-input-response":
				outputMessage.WriteString(" The response parameter is invalid or malformed;")
			case "bad-request":
				outputMessage.WriteString(" The request is invalid or malformed;")
			case "timeout-or-duplicate":
				outputMessage.WriteString(" The response is no longer valid: either is too old or has been used previously;")
			default:
				outputMessage.WriteString(" Unknown error code;")
			}
		}
	} else {
		outputMessage.WriteString(" Unable to identify codes.")
	}
	return outputMessage.String()
}
