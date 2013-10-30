package templating

import (
	"bones/validation"
)

type BaseContext struct {
	TemplateName string
	Errors       []string
	Notices      []string
}

func NewBaseContext(templateName string) *BaseContext {
	return &BaseContext{TemplateName: templateName}
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
