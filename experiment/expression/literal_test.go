package expression

import (
	"reflect"
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
	3.14,
	true,
	"haha",
	map[string]interface{}{"abc": 1, "def": "world"},
}

func TestLiteral_Value(t *testing.T) {
	for _, test := range TestLiteralValueTable {
		literal := Literal{test}
		actual, err := literal.Value(&base.Context{})
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(test, actual) {
			t.Error(test, actual)
		}
	}
}
