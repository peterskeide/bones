package validation

import (
	"strings"
)

type ValidationError struct {
	Messages []string
}

func (e *ValidationError) Error() string {
	return strings.Join(e.Messages, ", ")
}

func (e *ValidationError) AddMessage(message string) {
	e.Messages = append(e.Messages, message)
}
