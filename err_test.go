package err

import "testing"

func TestErr(t *testing.T) {
	// err := &errs{}

	// err.Err("c1", "e1")
	// t.Log("setSolo:", err.Detail())

	// err.Err("c2", "e2")
	// t.Log("setSolo:", err.Detail())

	// err.Err("c3", "e3")
	// t.Log("setSolo:", err.Detail())

	err := New()
	t.Log("err.GetCode():", err.GetCode())
}
