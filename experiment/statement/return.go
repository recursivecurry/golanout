package statement

import (
	"github.com/recursivecurry/golanout/experiment/base"
	"github.com/recursivecurry/golanout/experiment/expression"
)

type ReturnError struct {}
func (r *ReturnError) Error() string {
	return "return error"
}

type Return struct {
	value expression.Interface
}

func (r Return) Run(inputs base.Inputs, params base.Params, salt base.Salt) (base.Params, error) {
	_, _ = r.value.Value(inputs, params, salt)
	return params, &ReturnError{}
}

func (Return) Name() base.Operator {
	return "return"
}
