package expression

import (
	"github.com/recursivecurry/golanout/experiment/base"
)

type Or struct {
	values [2]Interface
}

func (o Or) Value(ctx *base.Context) (base.Value, error) {
	b1, err := GetBool(ctx, o.values[0])
	if err != nil {
		return nil, err
	}
	b2, err := GetBool(ctx, o.values[1])
	if err != nil {
		return nil, err
	}
	return b1 || b2, nil
}

func (Or) Name() base.Operator {
	return "or"
}
