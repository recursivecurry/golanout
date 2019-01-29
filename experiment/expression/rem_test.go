package expression

import (
	"github.com/recursivecurry/golanout/experiment/base"
	"testing"
)

func TestRem_Name(t *testing.T) {
	rem := Rem{}

	if rem.Name() != "%" {
		t.Fail()
	}
}

var TestRemValueTable = []struct {
	values   [2]Interface
	expected base.Value
	err      error
}{
	{
		values:   [2]Interface{Literal{20.0}, Literal{9.0}},
		expected: 2.0,
	},
	{
		values:   [2]Interface{Literal{-5.0}, Literal{2.0}},
		expected: -1.0,
	},
}

func TestRem_Value(t *testing.T) {
	for _, test := range TestRemValueTable {
		rem := Rem{test.values}
		actual, err := rem.Value(&base.Context{})
		if err != nil {
			if err.Error() != test.err.Error() {
				t.Error(err)
			}
		}
		if test.expected != actual {
			t.Error(test.expected, actual)
		}
	}
}
