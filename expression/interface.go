package expression

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/recursivecurry/golanout/base"
)

// A type that satisfies the pure.Interface can be a base.Value
type Interface interface {
	base.Interface
	// Value return the pure value based on inputs, params and salt.
	Value(*base.Context) (base.Value, error)
}

// GetBool return a bool value from a pure.Interface value.
func GetBool(env *base.Context, value Interface) (bool, error) {
	v, err := value.Value(env)
	if err != nil {
		return false, err
	}
	switch v := v.(type) {
	case bool:
		return v, nil
	case base.Number:
		return v != 0, nil
	default:
		return false, fmt.Errorf("wrong bool value: %+v", v)
	}
}

// GetNumber return a number value from a pure.Interface value.
func GetNumber(env *base.Context, value Interface) (base.Number, error) {
	v, err := value.Value(env)
	if err != nil {
		return 0, err
	}
	switch v := v.(type) {
	case base.Number:
		return v, nil
	default:
		return 0, fmt.Errorf("wrong number value: %+v", v)
	}
}

func GetInt(env *base.Context, value Interface) (int, error) {
	v, err := value.Value(env)
	if err != nil {
		return 0, err
	}
	if fv, ok := v.(base.Number); ok && base.Number(int(fv)) == fv {
		return int(fv), nil
	}
	return 0, errors.Errorf("non int value %+v", v)
}
