package impure

import "github.com/recursivecurry/golanout/experiment/base"

type Seq struct {
	seq []Interface
}

func (s Seq) Run(inputs base.Inputs, params base.Params, salt base.Salt) (base.Params, error) {
	for _, op := range s.seq {
		p, err := op.Run(inputs, params, salt)
		if err != nil {
			return nil, err
		}
		for k, v := range p {
			params[k] = v
		}
	}
	return params, nil
}

func (Seq) Name() base.Operator {
	return "seq"
}
