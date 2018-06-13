package pure

import "testing"

func TestAnd_Name(t *testing.T) {
	and := And{}

	if and.Name() != "and" {
		t.Fail()
	}
}

var TestAndRunTable = []struct {
	left     bool
	right    bool
	expected bool
}{
	{true, true, true},
	{true, false, false},
	{false, true, false},
	{false, false, false},
}

func TestAnd_Run(t *testing.T) {
	for _, test := range TestAndRunTable {
		and := And{[2]Interface{Literal{test.left}, Literal{test.right}}}
		actual, err := and.Value(nil, nil, "")
		if v, ok := actual.(bool); !ok || err != nil || v != test.expected {
			t.Fail()
		}
	}
}
