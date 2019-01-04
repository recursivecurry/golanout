package experiment

import (
	"github.com/recursivecurry/golanout/experiment/base"
	"github.com/recursivecurry/golanout/experiment/interpreter"
	"github.com/recursivecurry/golanout/experiment/statement"
)

type BasicExperiment struct {
	Name      string
	Overrides map[string]interface{}
	Code      statement.Interface

	salt      string
	evaluated bool
}

func (e *BasicExperiment) Run(inputs base.Inputs, overrides base.Overrides) (base.Params, error) {
	params := make(base.Params)
	params, err := e.Code.Run(inputs, params, base.Salt(e.salt))
	if err != nil {
		return nil, err
	}
	return params, nil
}

func NewBasicExperiment(name string, jsCode map[string]interface{}) Experiment {
	code := interpreter.Interprete(jsCode)
	return &BasicExperiment{
		Name: name,
		Overrides: make(map[string]interface{}),
		Code: code,

		salt: name,
	}
}
