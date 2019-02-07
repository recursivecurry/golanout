package expression

import (
	"strings"
	"testing"

	"github.com/recursivecurry/golanout/base"
)

func TestGetBool(t *testing.T) {
	var testData = []struct {
		input    Interface
		expected bool
	}{
		{
			Literal{base.Number(3.14)},
			true,
		},
		{
			Literal{base.Number(0.0)},
			false,
		},
		{
			Literal{true},
			true,
		},
		{
			Literal{false},
			false,
		},
	}
	for i, d := range testData {
		if b, err := GetBool(&base.Context{}, d.input); err != nil || b != d.expected {
			t.Error(i, d.expected, b)
		}
	}
}

func TestGetNumber(t *testing.T) {
	var testData = []struct {
		input    Interface
		expected base.Number
	}{
		{
			Literal{base.Number(3.14)},
			3.14,
		},
		{
			Literal{base.Number(1.0)},
			1.0,
		},
		{
			Literal{base.Number(0)},
			0,
		},
	}
	for i, d := range testData {
		if b, err := GetNumber(&base.Context{}, d.input); err != nil || b != d.expected {
			t.Error(i, d.expected, b)
		}
	}
}

func TestGetInt(t *testing.T) {
	var testData = []struct {
		input    Interface
		expected int
		err      string
	}{
		{
			Literal{base.Number(3.14)},
			0,
			"non int value",
		},
		{
			Literal{"a"},
			0,
			"non int value",
		},
		{
			Literal{base.Number(1.0)},
			1,
			"",
		},
		{
			Literal{base.Number(-3.0)},
			-3,
			"",
		},
	}
	for i, d := range testData {
		if b, err := GetInt(&base.Context{}, d.input); (err != nil && !strings.HasPrefix(err.Error(), d.err)) || b != d.expected {
			t.Error(i, d.expected, b, err)
		}
	}
}
