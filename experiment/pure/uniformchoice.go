package pure

import (
	"github.com/recursivecurry/golanout/experiment/base"
)

type UniformChoice struct {
	Choices []base.Value
	Unit    pure.Interface
}

func (UniformChoice) Name() base.Operator {
	return "uniformChoice"
}

func (UniformChoice) Value(inputs base.Inputs, params base.Params, salt base.Salt) base.Value {
	panic("implement me")
}
