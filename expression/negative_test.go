package expression

import (
	"testing"

	"github.com/recursivecurry/golanout/base"
)

func TestNegative_Name(t *testing.T) {
	negative := Negative{}

	if negative.Name() != "negative" {
		t.Fail()
	}
}

var TestNegativeValueTable = []struct {
	value    base.Number
	expected base.Number
}{
	{0, 0},
	{1, -1},
	{-1, 1},
}

func TestNegative_Value(t *testing.T) {
	for _, test := range TestNegativeValueTable {
		negative := Negative{Literal{test.value}}
		actual, err := negative.Value(&base.Context{})
		if v, ok := actual.(base.Number); !ok || err != nil || v != test.expected {
			t.Fail()
		}
	}
}
