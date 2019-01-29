package expression

import (
	"github.com/recursivecurry/golanout/experiment/base"
)

type Not struct {
	value Interface
}

func (n Not) Value(ctx *base.Context) (base.Value, error) {
	b, err := GetBool(ctx, n.value)
	if err != nil {
		return nil, err
	}
	return !b, nil
}

func (Not) Name() base.Operator {
	return "not"
}
