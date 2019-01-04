package expression

import (
	"errors"
	"github.com/recursivecurry/golanout/experiment/base"
)

type Get struct {
	Variable string
}

func (g Get) Value(inputs base.Inputs, params base.Params, salt base.Salt) (base.Value, error) {
	if v, ok := inputs[g.Variable]; ok {
		return v, nil
	}

	if v, ok := params[g.Variable]; ok {
		return v, nil
	}

	return nil, errors.New("non-exist variable")
}

func (Get) Name() base.Operator {
	return "get"
}
