package arr

import "testing"

func TestSumInt(t *testing.T) {
	expected := 6
	result := SumInt(1, 2, 3)

	if expected != result {
		t.Errorf("Should have %d, instead: %d\n", expected, result)
	}
}
