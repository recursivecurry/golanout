package expression

import (
	"testing"

	"github.com/recursivecurry/golanout/experiment/base"
)

func TestLiteral_Name(t *testing.T) {
	literal := Literal{}

	if literal.Name() != "literal" {
		t.Fail()
	}
}

var TestLiteralValueTable = []base.Value{
	1,
	true,
	"haha",
}

func TestLiteral_Value(t *testing.T) {
	for _, test := range TestLiteralValueTable {
		literal := Literal{test}
		actual, err := literal.Value(&base.Context{})
		if err != nil {
			t.Error(err)
		}
		if test != actual {
			t.Error(test, actual)
		}
	}
}
