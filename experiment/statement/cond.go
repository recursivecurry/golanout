package statement

import (
	"github.com/recursivecurry/golanout/experiment/base"
	"github.com/recursivecurry/golanout/experiment/expression"
)

type Cond struct {
	cond []IfThen
}

type IfThen struct {
	If   expression.Interface
	Then Interface
}

func (c Cond) Execute(ctx *base.Context) error {
	for _, ifThen := range c.cond {
		b, err := expression.GetBool(ctx, ifThen.If)
		if err != nil {
			return err
		}
		if b {
			return ifThen.Then.Execute(ctx)
		}
	}
	return nil
}

func (Cond) Name() base.Operator {
	return "cond"
}
