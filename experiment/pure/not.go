package pure

import (
	"github.com/recursivecurry/golanout/experiment/base"
)

type Not struct {
	value Interface
}

func (n Not) Value(inputs base.Inputs, params base.Params, salt base.Salt) (base.Value, error) {
	b, err := GetBool(n.value, inputs, params, salt)
	if err != nil {
		return nil, err
	}
	return !b, nil
}

func (Not) Name() base.Operator {
	return "not"
}
