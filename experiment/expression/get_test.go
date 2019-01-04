package expression

import "testing"

func TestGet_Name(t *testing.T) {
	get := Get{}

	if get.Name() != "get" {
		t.Fail()
	}
}

func TestGet_Value(t *testing.T) {
	get := Get{"b"}
	inputs := map[string]interface{}{"a": true, "b": 1.0}
	params := map[string]interface{}{"b": false, "c": 2}
	v, err := get.Value(inputs, params, "")
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
	inputs = map[string]interface{}{"a": true}
	params = map[string]interface{}{"b": false, "c": 2}
	v2, err := get.Value(inputs, params, "")
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
