package expression

import (
	"github.com/recursivecurry/golanout/base"
)

type And struct {
	values [2]Interface
}

func (a And) Value(ctx *base.Context) (base.Value, error) {
	b1, err := GetBool(ctx, a.values[0])
	if err != nil {
		return nil, err
	}
	b2, err := GetBool(ctx, a.values[1])
	if err != nil {
		return nil, err
	}
	return b1 && b2, nil
}

func (And) Name() base.Operator {
	return "and"
}
