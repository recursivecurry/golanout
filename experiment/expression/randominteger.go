package expression

import "github.com/recursivecurry/golanout/experiment/base"

type RandomInteger struct {
	Random
	Min int
	Max int
}

func (RandomInteger) Name() base.Operator {
	return "randomInteger"
}

func (r *RandomInteger) Value(ctx *base.Context) (base.Value, error) {
	return Literal{r.Min + r.getHash() % (r.Max - r.Min + 1)}, nil
}
