package base

type Salt string
type Operator string
type Value interface{}
type Params map[string]interface{}
type Inputs map[string]interface{}

// A type that satisfies the base.Interface can be a PlanOut operator
type Interface interface {
	Name() Operator
}
