package expression

import (
	"reflect"
	"testing"

	"github.com/recursivecurry/golanout/experiment/base"
)

func TestArray_Name(t *testing.T) {
	array := Array{}

	if array.Name() != "array" {
		t.Fail()
	}
}

var TestArrayValueTable = []struct {
	values   []Interface
	expected interface{}
}{
	{
		[]Interface{Literal{1}, Literal{true}, Literal{"haha"}},
		[]base.Value{1, true, "haha"},
	},
}

func TestArray_Value(t *testing.T) {
	for _, test := range TestArrayValueTable {
		array := Array{test.values}
		actual, err := array.Value(&base.Context{})
		if err != nil {
			t.Error(err)
		}
		if v, ok := actual.([]base.Value); !ok || !reflect.DeepEqual(test.expected, v) {
			t.Error(ok, test.expected, v)
		}
	}
}
