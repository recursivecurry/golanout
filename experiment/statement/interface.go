package statement

import "github.com/recursivecurry/golanout/experiment/base"

// A type that satisfies impure.Interface can be some operation with effect.
type Interface interface {
	base.Interface
	// Run execute the impure(effectual) operation based on inputs, params and salt.
	Run(base.Inputs, base.Params, base.Salt) (base.Params, error)
}
