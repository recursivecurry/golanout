package expression

import "testing"

func TestNegative_Name(t *testing.T) {
	negative := Negative{}

	if negative.Name() != "negative" {
		t.Fail()
	}
}

var TestNegativeValueTable = []struct {
	value    float64
	expected float64
}{
	{0, 0},
	{1, -1},
	{-1, 1},
}

func TestNegative_Value(t *testing.T) {
	for _, test := range TestNegativeValueTable {
		negative := Negative{Literal{test.value}}
		actual, err := negative.Value(nil, nil, "")
		if v, ok := actual.(float64); !ok || err != nil || v != test.expected {
			t.Fail()
		}
	}
}
