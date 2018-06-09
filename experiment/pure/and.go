package pure

import (
	"github.com/recursivecurry/golanout/experiment/base"
)

type And struct {
	values [2]Interface
}

func (a And) Run(inputs base.Inputs, params base.Params, salt base.Salt) (base.Value, error) {
	b1, err := GetBool(a.values[0], inputs, params, salt)
	if err != nil {
		return nil, err
	}
	b2, err := GetBool(a.values[1], inputs, params, salt)
	if err != nil {
		return nil, err
	}
	return b1 && b2, nil
}

func (And) Name() base.Operator {
	return "and"
}
