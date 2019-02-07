package expression

import (
	"testing"

	"github.com/recursivecurry/golanout/experiment/base"
)

func TestGet_Name(t *testing.T) {
	get := Get{}

	if get.Name() != "get" {
		t.Fail()
	}
}

func TestGet_Value(t *testing.T) {
	get := Get{"b"}
	ctx := &base.Context{
		Variables: map[string]interface{}{"a": true, "b": 1.0, "c": 2},
	}
	v, err := get.Value(ctx)
	if err != nil {
		t.Fail()
	}
	n, ok := v.(float64)
	if !ok {
		t.Fail()
	}
	if n != 1 {
		t.Fail()
	}
	ctx = &base.Context{
		Variables: map[string]interface{}{"a": true, "b": false, "c": 2},
	}
	v2, err := get.Value(ctx)
	if err != nil {
		t.Fail()
	}
	n2, ok := v2.(bool)
	if !ok {
		t.Fail()
	}
	if n2 != false {
		t.Fail()
	}
}
