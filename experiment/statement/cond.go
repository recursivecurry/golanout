package statement

import (
	"github.com/recursivecurry/golanout/experiment/base"
	"github.com/recursivecurry/golanout/experiment/expression"
)

type Cond struct {
	cond []IfThen
}

type IfThen struct {
	If   expression.Interface
	Then Interface
}

func (c Cond) Run(inputs base.Inputs, params base.Params, salt base.Salt) (base.Params, error) {
	for _, ifThen := range c.cond {
		b, err := expression.GetBool(ifThen.If, inputs, params, salt)
		if err != nil {
			return params, err
		}
		if b {
			return ifThen.Then.Run(inputs, params, salt)
		}
	}
	return params, nil
}

func (Cond) Name() base.Operator {
	return "cond"
}
