package experiment

import (
	"github.com/recursivecurry/golanout/experiment/base"
	"github.com/recursivecurry/golanout/experiment/interpreter"
	"github.com/recursivecurry/golanout/experiment/statement"
)

type Experiment struct {
	Name      string
	Code      statement.Interface
	salt      string
}
func (e *Experiment) Execute(inputs map[string]interface{}, overrides map[string]interface{}) (base.Variables, error) {
	variables := make(base.Variables)
	for k, v := range inputs {
		variables[k] = v
	}
	for k, v := range overrides {
		variables[k] = v
	}
	env := &base.Context{
		Variables: variables,
		Salt: e.salt,
	}
	err := e.Code.Execute(env)
	if err != nil {
		return nil, err
	}
	return env.Variables, nil
}

func NewExperiment(name string, jsCode map[string]interface{}) Interface {
	code := interpreter.Interprete(jsCode)
	return &Experiment{
		Name: name,
		Code: code,
		salt: name,
	}
}
