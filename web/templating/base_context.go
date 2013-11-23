package templating

import (
	"bones/validation"
	"fmt"
	"html/template"
)

type BaseContext struct {
	TemplateName string
	Errors       []string
	Notices      []string
	CsrfToken    string
}

func (ctx *BaseContext) AddError(err error) {
	switch t := err.(type) {
	case *validation.ValidationError:
		for _, message := range t.Messages {
			ctx.Errors = append(ctx.Errors, message)
		}
	default:
		ctx.Errors = append(ctx.Errors, t.Error())
	}
}

func (ctx *BaseContext) AddNotice(notice string) {
	ctx.Notices = append(ctx.Notices, notice)
}

func (ctx *BaseContext) Name() string {
	return ctx.TemplateName
}

func (ctx *BaseContext) CsrfTokenField() template.HTML {
	inputField := fmt.Sprintf(`<input type="hidden" id="crsf-token" name="CsrfToken" value="%s">`, ctx.CsrfToken)
	return template.HTML(inputField)
}
