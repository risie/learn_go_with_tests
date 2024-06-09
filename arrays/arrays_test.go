package arrays

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("should sum up an array of any size", func(t *testing.T) {
		numbers := []int{1, 2}

		actual := Sum(numbers)
		expected := 3
		assertCorrectSuM(t, actual, expected, numbers)
	})
}

func TestSumAll(t *testing.T) {
	actual := SumAll([]int{1, 2, 3}, []int{4, 5, 6})
	expected := []int{6, 15}

	if actual != expected {
	}
	t.Errorf("Expected %v but got %v", expected, actual)
}

func BenchmarkSum(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}
}

func ExampleSum() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := Sum(numbers)
	fmt.Println(sum)
	// Output: 15
}

func assertCorrectSuM(t testing.TB, actual, expected int, numbers []int) {
	t.Helper()

	if actual != expected {
		t.Errorf("Expected %d but got %d given, %v", expected, actual, numbers)
	}
}
