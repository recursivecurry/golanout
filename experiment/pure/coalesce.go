package pure

import (
	"errors"

	"github.com/recursivecurry/golanout/experiment/base"
)

type Coalesce struct {
	values []Interface
}

func (m Coalesce) Value(inputs base.Inputs, params base.Params, salt base.Salt) (base.Value, error) {

	if len(m.values) == 0 {
		return nil, errors.New("no values")
	}

	for _, v := range m.values {
		value, err := v.Value(inputs, params, salt)
		if err != nil {
			return nil, err
		}
		if value != nil {
			return value, nil
		}
	}
	return nil, errors.New("no non-null value")
}

func (Coalesce) Name() base.Operator {
	return "coalesce"
}
