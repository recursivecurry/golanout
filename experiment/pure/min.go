package pure

import (
	"errors"
	"github.com/recursivecurry/golanout/experiment/base"
)

type Min struct {
	values []Interface
}

func (m Min) Value(inputs base.Inputs, params base.Params, salt base.Salt) (base.Value, error) {

	if len(m.values) == 0 {
		return nil, errors.New("no values")
	}
	var min float64
	value, err := m.values[0].Value(inputs, params, salt)
	if err != nil {
		return nil, err
	}
	if fv, ok := value.(float64); ok {
		min = fv
	}
	for _, v := range m.values[1:] {
		value, err := v.Value(inputs, params, salt)
		if err != nil {
			return nil, err
		}
		if fv, ok := value.(float64); ok {
			if fv < min {
				min = fv
			}
		}
	}
	return min, nil
}

func (Min) Name() base.Operator {
	return "min"
}
