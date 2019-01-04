package expression

import (
	"testing"

	"github.com/recursivecurry/golanout/experiment/base"
)

func TestMax_Name(t *testing.T) {
	max := Max{}

	if max.Name() != "max" {
		t.Fail()
	}
}

var TestMaxValueTable = []struct {
	values   []Interface
	expected base.Value
}{
	{
		values:   []Interface{Literal{10.0}, Literal{5.0}, Literal{20.0}},
		expected: 20.0,
	},
}

func TestMax_Value(t *testing.T) {
	for _, test := range TestMaxValueTable {
		max := Max{test.values}
		actual, err := max.Value(nil, nil, "")
		if err != nil {
			t.Error(err)
		}
		if test.expected != actual {
			t.Error(test.expected, actual)
		}
	}
}
