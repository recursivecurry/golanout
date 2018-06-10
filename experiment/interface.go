package experiment

import (
	"github.com/recursivecurry/golanout/experiment/base"
)

type Experiment interface {
	Run(base.Inputs) base.Params
}
