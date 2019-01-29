package expression

import (
	"github.com/recursivecurry/golanout/experiment/base"
)

type Sum struct {
	values [2]Interface
}

func (s Sum) Value(ctx *base.Context) (base.Value, error) {
	n1, err := GetNumber(ctx, s.values[0])
	if err != nil {
		return nil, err
	}
	n2, err := GetNumber(ctx, s.values[1])
	if err != nil {
		return nil, err
	}
	return n1 + n2, nil
}

func (Sum) Name() base.Operator {
	return "sum"
}
