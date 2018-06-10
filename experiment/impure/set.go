package impure

import (
	"github.com/recursivecurry/golanout/experiment/base"
	"github.com/recursivecurry/golanout/experiment/pure"
)

type Set struct {
	variable string
	value    pure.Interface
}

func (s Set) Run(inputs base.Inputs, params base.Params, salt base.Salt) (base.Params, error) {
	v, err := s.value.Value(inputs, params, salt)
	if err != nil {
		return nil, err
	}
	params[s.variable] = v
	return params, nil
}

func (Set) Name() base.Operator {
	return "set"
}
