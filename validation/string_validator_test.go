package validation

import (
	"testing"
)

func Test_StringValidator_NotEmpty_addsMessageIfStringIsEmpty(t *testing.T) {
	v, err := setupStringValidator("")
	v.NotEmpty("should not be empty")
	assertErrorMatches(t, err, "should not be empty")
}

func Test_StringValidator_NotEmpty_addsNoMessageIfStringIsNotEmpty(t *testing.T) {
	v, err := setupStringValidator("test")
	v.NotEmpty("should not be empty")
	assertNoErrorMessages(t, err)
}

func Test_StringValidator_NotEmpty_supportsChaining(t *testing.T) {
	v, _ := setupStringValidator("test")
	retval := v.NotEmpty("should not be empty")
	assertSameStringValidator(t, v, retval)
}

func Test_StringValidator_Equals_addsMessageIfStringsAreNotEqual(t *testing.T) {
	v, err := setupStringValidator("test")
	v.Equals("foo", "should be equal")
	assertErrorMatches(t, err, "should be equal")
}

func Test_StringValidator_Equals_addsNoMessageIfStringsAreEqual(t *testing.T) {
	v, err := setupStringValidator("test")
	v.Equals("test", "should be equal")
	assertNoErrorMessages(t, err)
}

func Test_StringValidator_Equals_supportsChaining(t *testing.T) {
	v, _ := setupStringValidator("test")
	retval := v.Equals("test", "should not be empty")
	assertSameStringValidator(t, v, retval)
}

func Test_StringValidator_MaxLength_addsMessageIfStringIsLongerThanMaxLength(t *testing.T) {
	v, err := setupStringValidator("test")
	v.MaxLength(3, "field must be lt or eq 3")
	assertErrorMatches(t, err, "field must be lt or eq 3")
}

func Test_StringValidator_MaxLength_addsNoMessageIfStringIsShorterThanOrEqualToMaxLength(t *testing.T) {
	v, err := setupStringValidator("test")
	v.MaxLength(4, "field must be lt or eq 4")
	assertNoErrorMessages(t, err)
}

func Test_StringValidator_MaxLength_supportsChaining(t *testing.T) {
	v, _ := setupStringValidator("test")
	retval := v.MaxLength(4, "should not be empty")
	assertSameStringValidator(t, v, retval)
}

func Test_StringValidator_MinLength_addsMessageIfStringIsShorterThanMinLength(t *testing.T) {
	v, err := setupStringValidator("test")
	v.MinLength(5, "field must be gt or eq 5")
	assertErrorMatches(t, err, "field must be gt or eq 5")
}

func Test_StringValidator_MinLength_addsNoMessageIfStringIsLongerThanOrEqualToMinLength(t *testing.T) {
	v, err := setupStringValidator("test")
	v.MinLength(4, "field must be gt or eq 4")
	assertNoErrorMessages(t, err)
}

func Test_StringValidator_MinLength_supportsChaining(t *testing.T) {
	v, _ := setupStringValidator("test")
	retval := v.MinLength(4, "should not be empty")
	assertSameStringValidator(t, v, retval)
}

func setupStringValidator(value string) (*StringValidator, *ValidationError) {
	err := new(ValidationError)
	v := &StringValidator{value, err}
	return v, err
}

func assertErrorMatches(t *testing.T, err error, expected string) {
	if err.Error() != expected {
		t.Error("Expected validation to fail")
	}
}

func assertNoErrorMessages(t *testing.T, err *ValidationError) {
	if len(err.Messages) > 0 {
		t.Error("Expected validation to succeed")
	}
}

func assertSameStringValidator(t *testing.T, expected *StringValidator, actual *StringValidator) {
	if expected != actual {
		t.Error("Expected method to support chaining")
	}
}
