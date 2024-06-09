package arrays

import "testing"

func TestSum(t *testing.T) {
	t.Run("should sum up an array of 5 integers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		actual := Sum(numbers)
		expected := 15
		assertCorrectSuM(t, actual, expected, numbers)
	})
	t.Run("should sum up an array of any size", func(t *testing.T) {
		numbers := []int{1, 2}

		actual := Sum(numbers)
		expected := 3
		assertCorrectSuM(t, actual, expected, numbers)
	})
}

func BenchmarkSum(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}
}

func assertCorrectSuM(t testing.TB, actual, expected int, numbers []int) {
	t.Helper()

	if actual != expected {
		t.Errorf("Expected %d but got %d given, %v", expected, actual, numbers)
	}
}
