package expression

import (
	"github.com/recursivecurry/golanout/experiment/base"
	"testing"
)

func TestNot_Name(t *testing.T) {
	not := Not{}

	if not.Name() != "not" {
		t.Fail()
	}
}

var TestNotBoolValueTable = []struct {
	value    bool
	expected bool
}{
	{true, false},
	{false, true},
}

func TestNotBool_Value(t *testing.T) {
	for _, test := range TestNotBoolValueTable {
		not := Not{Literal{test.value}}
		actual, err := not.Value(&base.Context{})
		if v, ok := actual.(bool); !ok || err != nil || v != test.expected {
			t.Fail()
		}
	}
}

var TestNotNumberValueTable = []struct {
	value    float64
	expected bool
}{
	{0, true},
	{1, false},
}

func TestNot_Value(t *testing.T) {
	for _, test := range TestNotNumberValueTable {
		not := Not{Literal{test.value}}
		actual, err := not.Value(&base.Context{})
		if v, ok := actual.(bool); !ok || err != nil || v != test.expected {
			t.Fail()
		}
	}
}
