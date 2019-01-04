package expression

import (
	"github.com/recursivecurry/golanout/experiment/base"
)

type Sum struct {
	values [2]Interface
}

func (s Sum) Value(inputs base.Inputs, params base.Params, salt base.Salt) (base.Value, error) {
	n1, err := GetNumber(s.values[0], inputs, params, salt)
	if err != nil {
		return nil, err
	}
	n2, err := GetNumber(s.values[1], inputs, params, salt)
	if err != nil {
		return nil, err
	}
	return n1 + n2, nil
}

func (Sum) Name() base.Operator {
	return "sum"
}
