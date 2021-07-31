package str

import "testing"

func TestToUpper(t *testing.T) {
	expected := "HELLO"
	result := ToUpper("hello")

	if result != expected {
		t.Errorf("Should have %s, instead: %s\n", expected, result)
	}
}
