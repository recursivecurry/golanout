package pure

import (
	"reflect"
	"testing"

	"github.com/recursivecurry/golanout/experiment/base"
)

func TestArray_Name(t *testing.T) {
	and := Array{}

	if and.Name() != "array" {
		t.Fail()
	}
}

var TestArrayRunTable = []struct {
	values   []Interface
	expected interface{}
}{
	{
		[]Interface{Literal{1}, Literal{true}, Literal{"haha"}},
		[]base.Value{1, true, "haha"},
	},
}

func TestArray_Run(t *testing.T) {
	for _, test := range TestArrayRunTable {
		array := Array{test.values}
		actual, err := array.Value(nil, nil, "")
		if err != nil {
			t.Error(err)
		}
		if v, ok := actual.([]base.Value); !ok || !reflect.DeepEqual(test.expected, v) {
			t.Error(ok, test.expected, v)
		}
	}
}
