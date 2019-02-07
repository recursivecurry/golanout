package expression

import (
	"errors"

	"github.com/recursivecurry/golanout/base"
)

type Max struct {
	values []Interface
}

func (m Max) Value(ctx *base.Context) (base.Value, error) {

	if len(m.values) == 0 {
		return nil, errors.New("no values")
	}
	var max base.Number
	value, err := m.values[0].Value(ctx)
	if err != nil {
		return nil, err
	}
	if fv, ok := value.(base.Number); ok {
		max = fv
	}
	for _, v := range m.values[1:] {
		value, err := v.Value(ctx)
		if err != nil {
			return nil, err
		}
		if fv, ok := value.(base.Number); ok {
			if fv > max {
				max = fv
			}
		}
	}
	return max, nil
}

func (Max) Name() base.Operator {
	return "max"
}
