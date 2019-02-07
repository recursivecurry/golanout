package expression

import (
	"testing"

	"github.com/recursivecurry/golanout/base"
)

func TestSum_Name(t *testing.T) {
	sum := Sum{}

	if sum.Name() != "sum" {
		t.Fail()
	}
}

var TestSumValueTable = []struct {
	values   [2]Interface
	expected base.Value
}{
	{
		values:   [2]Interface{Literal{10.0}, Literal{20.0}},
		expected: 30.0,
	},
	{
		values:   [2]Interface{Literal{-10.0}, Literal{5.0}},
		expected: -5.0,
	},
	{
		values:   [2]Interface{Literal{-10.0}, Literal{0.0}},
		expected: -10.0,
	},
}

func TestSum_Value(t *testing.T) {
	for _, test := range TestSumValueTable {
		sum := Sum{test.values}
		actual, err := sum.Value(&base.Context{})
		if err != nil {
			t.Error(err)
		}
		if test.expected != actual {
			t.Error(test.expected, actual)
		}
	}
}
