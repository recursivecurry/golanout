package impure

import (
	"github.com/recursivecurry/golanout/experiment/base"
	"github.com/recursivecurry/golanout/experiment/pure"
)

type Cond struct {
	cond []IfThen
}

type IfThen struct {
	If   pure.Interface
	Then Interface
}

func (c Cond) Run(inputs base.Inputs, params base.Params, salt base.Salt) (base.Params, error) {
	for _, ifThen := range c.cond {
		b, err := pure.GetBool(ifThen.If, inputs, params, salt)
		if err != nil {
			return nil, err
		}
		if b {
			return ifThen.Then.Run(inputs, params, salt)
		}
	}
	return params, nil
}

func (Cond) Name() base.Operator {
	return "seq"
}
