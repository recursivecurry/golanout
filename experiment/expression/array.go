package expression

import "github.com/recursivecurry/golanout/experiment/base"

type Array struct {
	Values []Interface
}

func (a Array) Value(ctx *base.Context) (base.Value, error) {
	var values = make([]base.Value, 0, len(a.Values))
	for _, v := range a.Values {
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
