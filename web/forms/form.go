package forms

type Form interface {
	Validate() error
}
