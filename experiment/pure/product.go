package pure

import (
	"github.com/recursivecurry/golanout/experiment/base"
)

type Product struct {
	values [2]Interface
}

func (p Product) Run(inputs base.Inputs, params base.Params, salt base.Salt) (base.Value, error) {
	n1, err := GetNumber(p.values[0], inputs, params, salt)
	if err != nil {
		return nil, err
	}
	n2, err := GetNumber(p.values[1], inputs, params, salt)
	if err != nil {
		return nil, err
	}
	return n1 * n2, nil
}

func (Product) Name() base.Operator {
	return "product"
}
