package expression

import (
	"errors"

	"github.com/recursivecurry/golanout/experiment/base"
)

type Get struct {
	Variable string
}

func (g Get) Value(ctx *base.Context) (base.Value, error) {
	if v, ok := ctx.Variables[g.Variable]; ok {
		return v, nil
	}

	return nil, errors.New("non-exist variable")
}

func (Get) Name() base.Operator {
	return "get"
}
