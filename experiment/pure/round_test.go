package pure

import "testing"

func TestRound_Name(t *testing.T) {
	round := Round{}

	if round.Name() != "round" {
		t.Fail()
	}
}

var TestRoundValueTable = []struct {
	value    float64
	expected float64
}{
	{0, 0},
	{1.3, 1.0},
	{1.5, 2.0},
	{-1.3, -1.0},
	{-1.5, -2.0},
}

func TestRound_Value(t *testing.T) {
	for _, test := range TestRoundValueTable {
		round := Round{Literal{test.value}}
		actual, err := round.Value(nil, nil, "")
		if v, ok := actual.(float64); !ok || err != nil || v != test.expected {
			t.Fail()
		}
	}
}
