package expression

import (
	"github.com/recursivecurry/golanout/experiment/base"
	"testing"
)

func TestLength_Name(t *testing.T) {
	length := Length{}

	if length.Name() != "length" {
		t.Fail()
	}
}

var TestLengthValueTable = []struct {
	value   Interface
	expected int
}{
	{
		value: Array{values: []Interface{Literal{1}, Literal{"a"}, Literal{3.14}}},
		expected: 3,
	},
	{
		value: Literal{value: map[string]interface{}{"a": 1, "b": 2}},
		expected: 2,
	},
	{
		value: Literal{value: "abcd"},
		expected: 4,
	},
}

func TestLength_Value(t *testing.T) {
	for _, test := range TestLengthValueTable {
		length := Length{test.value}
		actual, err := length.Value(&base.Context{})
		if err != nil {
			t.Error(err)
		}
		if v, ok := actual.(int); !ok || v != test.expected {
			t.Error(ok, test.expected, v)
		}
	}
}
