package exampletest

import "testing"

func TestSum(t *testing.T) {
	result := Sum(1, 1)
	want := float64(2)

	if result != want {
		t.Errorf("expected resutl tobe %v, got %v", result, want)
	}
}
