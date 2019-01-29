package statement

import "github.com/recursivecurry/golanout/experiment/base"

type Seq struct {
	seq []Interface
}

func (s Seq) Execute(env *base.Context) error {
	for _, op := range s.seq {
		err := op.Execute(env)
		if err != nil {
			return err
		}
	}
	return nil
}

func (Seq) Name() base.Operator {
	return "seq"
}
