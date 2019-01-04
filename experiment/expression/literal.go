package expression

import "github.com/recursivecurry/golanout/experiment/base"

type Literal struct {
	value base.Value
}

func (l Literal) Value(inputs base.Inputs, params base.Params, salt base.Salt) (base.Value, error) {
	return l.value, nil
}

func (Literal) Name() base.Operator {
	return "literal"
}
