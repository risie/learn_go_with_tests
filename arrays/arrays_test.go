package arrays

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	actual := Sum(numbers)
	expected := 15

	if actual != expected {
		t.Errorf("Expected %d but got %d given, %v", expected, actual, numbers)
	}
}
