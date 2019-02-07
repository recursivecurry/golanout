package expression

import (
	"errors"

	"github.com/recursivecurry/golanout/experiment/base"
)

type Quot struct {
	values [2]Interface
}

func (q Quot) Value(ctx *base.Context) (base.Value, error) {
	n1, err := GetNumber(ctx, q.values[0])
	if err != nil {
		return nil, err
	}
	n2, err := GetNumber(ctx, q.values[1])
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
