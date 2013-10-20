package validation

type Validator struct {
	err *ValidationError
}

func New() *Validator {
	return &Validator{new(ValidationError)}
}

func (v *Validator) Result() error {
	if len(v.err.Messages) > 0 {
		return v.err
	}

	return nil
}

func (v *Validator) String(value string) *StringValidator {
	return &StringValidator{value, v.err}
}
