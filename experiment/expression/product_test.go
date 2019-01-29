package expression

import (
	"github.com/recursivecurry/golanout/experiment/base"
	"testing"
)

func TestProduct_Name(t *testing.T) {
	product := Product{}

	if product.Name() != "product" {
		t.Fail()
	}
}

var TestProductValueTable = []struct {
	values   [2]Interface
	expected base.Value
}{
	{
		values:   [2]Interface{Literal{10.0}, Literal{20.0}},
		expected: 200.0,
	},
	{
		values:   [2]Interface{Literal{-10.0}, Literal{5.0}},
		expected: -50.0,
	},
	{
		values:   [2]Interface{Literal{-10.0}, Literal{0.0}},
		expected: 0.0,
	},
}

func TestProduct_Value(t *testing.T) {
	for _, test := range TestProductValueTable {
		product := Product{test.values}
		actual, err := product.Value(&base.Context{})
		if err != nil {
			t.Error(err)
		}
		if test.expected != actual {
			t.Error(test.expected, actual)
		}
	}
}
