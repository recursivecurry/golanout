package pure

import (
	"testing"
	"github.com/recursivecurry/golanout/experiment/base"
)

func TestMin_Name(t *testing.T) {
	min := Min{}

	if min.Name() != "min" {
		t.Fail()
	}
}

var TestMinRunTable = []struct{
	values []Interface
	expected base.Value
}{
	{
		values: []Interface{Literal{10.0}, Literal{5.0}, Literal{20.0}},
		expected: 5.0,
	},
}

func TestMin_Run(t *testing.T) {
	for _, test := range TestMinRunTable {
		min := Min{test.values}
		actual, err := min.Value(nil, nil, "")
		if err != nil {
			t.Error(err)
		}
		if test.expected != actual {
			t.Error(test.expected, actual)
		}
	}
}
