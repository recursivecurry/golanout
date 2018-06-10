package pure

import (
	"github.com/recursivecurry/golanout/experiment/base"
)

type Negative struct {
	value Interface
}

func (n Negative) Run(inputs base.Inputs, params base.Params, salt base.Salt) (base.Value, error) {
	number, err := GetNumber(n.value, inputs, params, salt)
	if err != nil {
		return nil, err
	}
	return -number, nil
}

func (Negative) Name() base.Operator {
	return "negative"
}
