package statement

import (
	"github.com/recursivecurry/golanout/experiment/base"
	"github.com/recursivecurry/golanout/experiment/expression"
)

type Set struct {
	variable string
	value    expression.Interface
}

func (s Set) Run(inputs base.Inputs, params base.Params, salt base.Salt) (base.Params, error) {
	v, err := s.value.Value(inputs, params, salt)
	if err != nil {
		return params, err
	}
	params[s.variable] = v
	return params, nil
}

func (Set) Name() base.Operator {
	return "set"
}
