package expression

import (
	"testing"

	"github.com/recursivecurry/golanout/base"
)

func TestAnd_Name(t *testing.T) {
	and := And{}

	if and.Name() != "and" {
		t.Fail()
	}
}

var TestAndValueTable = []struct {
	left     bool
	right    bool
	expected bool
}{
	{true, true, true},
	{true, false, false},
	{false, true, false},
	{false, false, false},
}

func TestAnd_Value(t *testing.T) {
	for _, test := range TestAndValueTable {
		and := And{[2]Interface{Literal{test.left}, Literal{test.right}}}
		actual, err := and.Value(&base.Context{})
		if v, ok := actual.(bool); !ok || err != nil || v != test.expected {
			t.Fail()
		}
	}
}

func TestAnd_WrongValue(t *testing.T) {
	and := And{[2]Interface{Literal{true}, Literal{1}}}
	_, err := and.Value(&base.Context{})
	if err == nil {
		t.Fail()
	}
}
