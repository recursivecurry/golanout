package expression

import (
	"github.com/recursivecurry/golanout/base"
)

type Product struct {
	values [2]Interface
}

func (p Product) Value(ctx *base.Context) (base.Value, error) {
	n1, err := GetNumber(ctx, p.values[0])
	if err != nil {
		return nil, err
	}
	n2, err := GetNumber(ctx, p.values[1])
	if err != nil {
		return nil, err
	}
	return n1 * n2, nil
}

func (Product) Name() base.Operator {
	return "product"
}
