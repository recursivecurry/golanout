package pure

import (
	"github.com/recursivecurry/golanout/experiment/base"
)

type Or struct {
	values [2]Interface
}

func (o Or) Run(inputs base.Inputs, params base.Params, salt base.Salt) (base.Value, error) {
	b1, err := GetBool(o.values[0], inputs, params, salt)
	if err != nil {
		return nil, err
	}
	b2, err := GetBool(o.values[1], inputs, params, salt)
	if err != nil {
		return nil, err
	}
	return b1 || b2, nil
}

func (Or) Name() base.Operator {
	return "or"
}