package experiment

import "github.com/recursivecurry/golanout/experiment/base"

type Interface interface {
	Execute(inputs map[string]interface{}, overrides map[string]interface{}) (base.Variables, error)
}