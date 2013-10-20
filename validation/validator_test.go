package validation

import (
	"testing"
)

func Test_Validator_String_returnsStringValidatorForTheGivenValue(t *testing.T) {
	validator := New()
	value := "foobar"

	stringValidator := validator.String(value)

	if stringValidator.value != value {
		t.Error("Expected stringValidator to contain given value")
	}

	if stringValidator.err != validator.err {
		t.Error("Expected stringValidator to have a reference to the ValidationError of the validator")
	}
}
