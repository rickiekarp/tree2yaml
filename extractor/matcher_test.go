package extractor

import (
	"testing"
)

func TestGetOccurrenceProbabilityMap(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		slice := []string{}
		expected := map[int]int{}
		actual := getOccurrenceProbabilityMap(slice)

		if len(actual) != len(expected) {
			t.Fatalf("Expected len %d but got %d", len(expected), len(actual))
		}
	})

	t.Run("slice with single element", func(t *testing.T) {
		slice := []string{"apple"}
		expected := map[int]int{1: 100}
		actual := getOccurrenceProbabilityMap(slice)

		if len(actual) != len(expected) {
			t.Fatalf("Expected len %d but got %d", len(expected), len(actual))
		}

		for i, val := range expected {
			if actual[i] != val {
				t.Fatalf("Expected %d but got %d", val, actual[i])
			}
		}
	})

	t.Run("slice with multiple elements", func(t *testing.T) {
		slice := []string{"apple", "banana", "cherry"}
		expected := map[int]int{1: 33, 2: 66, 3: 100}
		actual := getOccurrenceProbabilityMap(slice)

		if len(actual) != len(expected) {
			t.Fatalf("Expected len %d but got %d", len(expected), len(actual))
		}

		for i, val := range expected {
			if actual[i] != val {
				t.Fatalf("Expected %d but got %d", val, actual[i])
			}
		}
	})
}
