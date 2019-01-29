package expression

import "github.com/recursivecurry/golanout/experiment/base"

type RandomFloat struct {
	Random
	Min float64
	Max float64
}

func (RandomFloat) Name() base.Operator {
	return "randomFloat"
}

func (r *RandomFloat) Value(ctx *base.Context) (base.Value, error) {
	return Literal{r.getUniform(r.Min, r.Max, r.Unit)}, nil
}
