package pure

import (
	"github.com/recursivecurry/golanout/experiment/base"
)

type Rem struct {
	values [2]Interface
}

func (r Rem) Run(inputs base.Inputs, params base.Params, salt base.Salt) (base.Value, error) {
	n1, err := GetNumber(r.values[0], inputs, params, salt)
	if err != nil {
		return nil, err
	}
	n2, err := GetNumber(r.values[1], inputs, params, salt)
	if err != nil {
		return nil, err
	}
	return int(n1) % int(n2), nil
}

func (Rem) Name() base.Operator {
	return "%"
}
