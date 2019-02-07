package statement

import (
	"github.com/recursivecurry/golanout/base"
	"github.com/recursivecurry/golanout/expression"
)

type ReturnError struct{}

func (r *ReturnError) Error() string {
	return "return error"
}

type Return struct {
	value expression.Interface
}

func (r Return) Execute(ctx *base.Context) error {
	if logging, err := expression.GetBool(ctx, r.value); err != nil && logging {
		ctx.NoLog = false
	}
	return &ReturnError{}
}

func (Return) Name() base.Operator {
	return "return"
}
