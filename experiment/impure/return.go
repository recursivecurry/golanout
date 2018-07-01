package impure

import (
	"github.com/recursivecurry/golanout/experiment/base"
	"github.com/recursivecurry/golanout/experiment/pure"
)

type Return struct {
	value pure.Interface
}

func (r Return) Run(inputs base.Inputs, params base.Params, salt base.Salt) (base.Params, error) {
	_, err := r.value.Value(inputs, params, salt)
	if err != nil {
		return nil, err
	}
	return params, nil
}

func (Return) Name() base.Operator {
	return "return"
}
