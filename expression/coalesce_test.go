package expression

import (
	"testing"

	"github.com/recursivecurry/golanout/base"
)

func TestCoalesce_Name(t *testing.T) {
	coalesce := Coalesce{}

	if coalesce.Name() != "coalesce" {
		t.Fail()
	}
}

var TestCoalesceValueTable = []struct {
	values   []Interface
	expected base.Value
}{
	{
		values:   []Interface{Literal{nil}, Literal{5.0}, Literal{20.0}},
		expected: 5.0,
	},
}

func TestCoalesce_Value(t *testing.T) {
	for _, test := range TestCoalesceValueTable {
		coalesce := Coalesce{test.values}
		actual, err := coalesce.Value(&base.Context{})
		if err != nil {
			t.Error(err)
		}
		if test.expected != actual {
			t.Error(test.expected, actual)
		}
	}
}
