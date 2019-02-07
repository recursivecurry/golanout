package expression

import (
	"github.com/recursivecurry/golanout/base"
)

type Negative struct {
	value Interface
}

func (n Negative) Value(ctx *base.Context) (base.Value, error) {
	number, err := GetNumber(ctx, n.value)
	if err != nil {
		return nil, err
	}
	return -number, nil
}

func (Negative) Name() base.Operator {
	return "negative"
}
