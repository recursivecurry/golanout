package base

type Salt string
type Operator string
type Value interface{}
type Variables map[string]interface{}

const (
	VarVersion = "result"
	VarFullSalt = "full_salt"
	VarSalt = "salt"
)

func (v Variables) Result() string {
	if r, ok := v[VarVersion].(string); ok {
		return r
	}
	return ""
}


// Context is an experiment execution environment.
type Context struct {
	Variables Variables
	Salt string
	NoLog bool
}

// A type that satisfies the base.Interface can be a PlanOut operator
type Interface interface {
	Name() Operator
}
