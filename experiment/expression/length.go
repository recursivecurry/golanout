package expression

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/recursivecurry/golanout/experiment/base"
)

type Length struct {
	value Interface
}

func (l Length) Value(ctx *base.Context) (base.Value, error) {
	v, err := l.value.Value(ctx)
	if err != nil {
		return nil, err
	}
	rv := reflect.ValueOf(v)

	switch rv.Kind() {
	case reflect.Array, reflect.Slice, reflect.Map, reflect.String:
		return rv.Len(), nil
	}
	return nil, errors.WithMessagef(err, "uncountable value %+v", v)
}

func (Length) Name() base.Operator {
	return "length"
}
