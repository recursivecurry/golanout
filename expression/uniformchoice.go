package expression

import (
	"github.com/recursivecurry/golanout/base"
)

type UniformChoice struct {
	Choices []base.Value
	Unit    Interface
}

func (UniformChoice) Name() base.Operator {
	return "uniformChoice"
}

func (UniformChoice) Value(ctx *base.Context) (base.Value, error) {
	panic("implement me")
}
