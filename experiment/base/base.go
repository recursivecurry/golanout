package base

type Salt string
type Operator string
type Value interface{}
type Params map[string]interface{}
type Inputs map[string]interface{}

type Interface interface {
	Name() Operator
}
