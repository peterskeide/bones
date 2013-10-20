package validation

type StringValidator struct {
	value string
	err   *ValidationError
}

func (v *StringValidator) NotEmpty(message string) *StringValidator {
	if v.value == "" {
		v.err.AddMessage(message)
	}

	return v
}

func (v *StringValidator) Equals(other string, message string) *StringValidator {
	if v.value != other {
		v.err.AddMessage(message)
	}

	return v
}

func (v *StringValidator) MaxLength(length int, message string) *StringValidator {
	if len(v.value) > length {
		v.err.AddMessage(message)
	}

	return v
}

func (v *StringValidator) MinLength(length int, message string) *StringValidator {
	if len(v.value) < length {
		v.err.AddMessage(message)
	}

	return v
}
