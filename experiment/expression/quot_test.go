package expression

import (
	"errors"
	"github.com/recursivecurry/golanout/experiment/base"
	"testing"
)

func TestQuot_Name(t *testing.T) {
	quot := Quot{}

	if quot.Name() != "/" {
		t.Fail()
	}
}

var TestQuotValueTable = []struct {
	values   [2]Interface
	expected base.Value
	err      error
}{
	{
		values:   [2]Interface{Literal{20.0}, Literal{10.0}},
		expected: 2.0,
	},
	{
		values:   [2]Interface{Literal{-10.0}, Literal{5.0}},
		expected: -2.0,
	},
	{
		values: [2]Interface{Literal{-10.0}, Literal{0.0}},
		err:    errors.New("division by zero"),
	},
}

func TestQuot_Value(t *testing.T) {
	for _, test := range TestQuotValueTable {
		quot := Quot{test.values}
		actual, err := quot.Value(nil, nil, "")
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
