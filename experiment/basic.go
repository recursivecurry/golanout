package experiment

import (
	"github.com/recursivecurry/golanout/experiment/base"
	"github.com/recursivecurry/golanout/experiment/impure"
)

type BasicExperiment struct {
	Name      string
	Salt      string
	Overrides map[string]interface{}
	Code      impure.Interface
}

func (e *BasicExperiment) Run(inputs base.Inputs) (base.Params, error) {
	params := make(base.Params)
	params, err := e.Code.Run(inputs, params, base.Salt(e.Salt))
	if err != nil {
		return nil, err
	}
	return params, nil
}

func NewBasicExperiment(jsCode map[string]interface{}) Experiment {
	panic("xxx")
}
