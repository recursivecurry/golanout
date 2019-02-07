package expression

import "github.com/recursivecurry/golanout/base"

type Array struct {
	values []Interface
}

func (a Array) Value(ctx *base.Context) (base.Value, error) {
	var values = make([]base.Value, 0, len(a.values))
	for _, v := range a.values {
		value, err := v.Value(ctx)
		if err != nil {
			return nil, err
		}
		values = append(values, value)
	}
	return values, nil
}

func (Array) Name() base.Operator {
	return "array"
}
