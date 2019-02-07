package statement

import "github.com/recursivecurry/golanout/base"

// A type that satisfies impure.Interface can be some operation with effect.
type Interface interface {
	base.Interface
	// Execute execute the impure(effectual) operation based on inputs, params and salt.
	Execute(*base.Context) error
}
