package expression

import "testing"

func TestOr_Name(t *testing.T) {
	or := Or{}

	if or.Name() != "or" {
		t.Fail()
	}
}

var TestOrValueTable = []struct {
	left     bool
	right    bool
	expected bool
}{
	{true, true, true},
	{true, false, true},
	{false, true, true},
	{false, false, false},
}

func TestOr_Value(t *testing.T) {
	for _, test := range TestOrValueTable {
		or := Or{[2]Interface{Literal{test.left}, Literal{test.right}}}
		actual, err := or.Value(nil, nil, "")
		if v, ok := actual.(bool); !ok || err != nil || v != test.expected {
			t.Fail()
		}
	}
}
