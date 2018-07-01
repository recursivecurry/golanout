package pure

import (
	"errors"
	"github.com/recursivecurry/golanout/experiment/base"
)

type Quot struct {
	values [2]Interface
}

func (q Quot) Value(inputs base.Inputs, params base.Params, salt base.Salt) (base.Value, error) {
	n1, err := GetNumber(q.values[0], inputs, params, salt)
	if err != nil {
		return nil, err
	}
	n2, err := GetNumber(q.values[1], inputs, params, salt)
	if err != nil {
		return nil, err
	}
	if n2 == 0 {
		return nil, errors.New("division by zero")
	}
	return n1 / n2, nil
}

func (Quot) Name() base.Operator {
	return "/"
}
