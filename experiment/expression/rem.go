package expression

import (
	"math"

	"github.com/recursivecurry/golanout/experiment/base"
)

type Rem struct {
	values [2]Interface
}

func (r Rem) Value(ctx *base.Context) (base.Value, error) {
	n1, err := GetNumber(ctx, r.values[0])
	if err != nil {
		return nil, err
	}
	n2, err := GetNumber(ctx, r.values[1])
	if err != nil {
		return nil, err
	}
	return math.Remainder(n1, n2), nil
}

func (Rem) Name() base.Operator {
	return "%"
}
