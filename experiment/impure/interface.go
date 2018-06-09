package impure

import "github.com/recursivecurry/golanout/experiment/base"

type Interface interface {
	base.Interface
	Run(base.Inputs, base.Params, base.Salt) (base.Params, error)
}
