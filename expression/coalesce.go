package expression

import (
	"errors"

	"github.com/recursivecurry/golanout/base"
)

type Coalesce struct {
	values []Interface
}

func (c Coalesce) Value(ctx *base.Context) (base.Value, error) {

	if len(c.values) == 0 {
		return nil, errors.New("no values")
	}

	for _, v := range c.values {
		value, err := v.Value(ctx)
		if err != nil {
			return nil, err
		}
		if value != nil {
			return value, nil
		}
	}
	return nil, errors.New("no non-null value")
}

func (Coalesce) Name() base.Operator {
	return "coalesce"
}
