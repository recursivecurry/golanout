package pure

import (
	"fmt"
	"github.com/recursivecurry/golanout/experiment/base"
)

type Interface interface {
	base.Interface
	Value(base.Inputs, base.Params, base.Salt) (base.Value, error)
}

func GetBool(value Interface, inputs base.Inputs, params base.Params, salt base.Salt) (bool, error) {
	v, err := value.Value(inputs, params, salt)
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

func GetNumber(value Interface, inputs base.Inputs, params base.Params, salt base.Salt) (float64, error) {
	v, err := value.Value(inputs, params, salt)
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
