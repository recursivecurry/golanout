package expression

import (
	"fmt"

	"github.com/recursivecurry/golanout/experiment/base"
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
	case float64:
		return v != 0, nil
	default:
		return false, fmt.Errorf("wrong bool value: %+v", v)
	}
}

// GetNumber return a number value from a pure.Interface value.
func GetNumber(env *base.Context, value Interface) (float64, error) {
	v, err := value.Value(env)
	if err != nil {
		return 0, err
	}
	switch v := v.(type) {
	case float64:
		return v, nil
	default:
		return 0, fmt.Errorf("wrong number value: %+v", v)
	}
}
