package expression

import (
	"github.com/recursivecurry/golanout/experiment/base"
	"math"
)

type Round struct {
	value Interface
}

func (r Round) Value(ctx *base.Context) (base.Value, error) {
	number, err := GetNumber(ctx, r.value)
	if err != nil {
		return nil, err
	}
	return math.Round(number), nil
}

func (Round) Name() base.Operator {
	return "round"
}
