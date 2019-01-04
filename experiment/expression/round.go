package expression

import (
	"github.com/recursivecurry/golanout/experiment/base"
	"math"
)

type Round struct {
	value Interface
}

func (r Round) Value(inputs base.Inputs, params base.Params, salt base.Salt) (base.Value, error) {
	number, err := GetNumber(r.value, inputs, params, salt)
	if err != nil {
		return nil, err
	}
	return math.Round(number), nil
}

func (Round) Name() base.Operator {
	return "round"
}
