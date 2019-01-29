package statement

import (
	"github.com/recursivecurry/golanout/experiment/base"
	"github.com/recursivecurry/golanout/experiment/expression"
)

type Set struct {
	variable string
	value    expression.Interface
}

func (s Set) Execute(env *base.Context) error {
	v, err := s.value.Value(env)
	if err != nil {
		return err
	}
	env.Variables[s.variable] = v
	return nil
}

func (Set) Name() base.Operator {
	return "set"
}
