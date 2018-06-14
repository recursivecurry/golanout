package pure

import (
	"testing"

	"github.com/recursivecurry/golanout/experiment/base"
)

func TestCoalesce_Name(t *testing.T) {
	and := Coalesce{}

	if and.Name() != "coalesce" {
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
		actual, err := coalesce.Value(nil, nil, "")
		if err != nil {
			t.Error(err)
		}
		if test.expected != actual {
			t.Error(test.expected, actual)
		}
	}
}
