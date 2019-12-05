package main

import (
	"reflect"
	"testing"
)

func TestIntcoder(t *testing.T) {
	t.Run("test process mutates self", func(t *testing.T) {
		testCases := [...]struct {
			input, expected Intcode
		}{
			{Intcode{1, 0, 0, 0, 99}, Intcode{2, 0, 0, 0, 99}},          // (1 + 1 = 2)
			{Intcode{2, 3, 0, 3, 99}, Intcode{2, 3, 0, 6, 99}},          // (3 * 2 = 6)
			{Intcode{2, 4, 4, 5, 99, 0}, Intcode{2, 4, 4, 5, 99, 9801}}, // (99 * 99 = 9801)
			{Intcode{1, 1, 1, 4, 99, 5, 6, 0, 99}, Intcode{30, 1, 1, 4, 2, 5, 6, 0, 99}},
		}

		for _, tt := range testCases {
			tt.input.process()
			if !reflect.DeepEqual(tt.input, tt.expected) {
				t.Errorf("got %v want %v", tt.input, tt.expected)
			}
		}
	})

	t.Run("test process returns the output at location 0", func(t *testing.T) {
		intcode := Intcode{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 13, 19, 1, 9, 19, 23, 1, 6, 23, 27, 2, 27, 9, 31, 2, 6, 31, 35, 1, 5, 35, 39, 1, 10, 39, 43, 1, 43, 13, 47, 1, 47, 9, 51, 1, 51, 9, 55, 1, 55, 9, 59, 2, 9, 59, 63, 2, 9, 63, 67, 1, 5, 67, 71, 2, 13, 71, 75, 1, 6, 75, 79, 1, 10, 79, 83, 2, 6, 83, 87, 1, 87, 5, 91, 1, 91, 9, 95, 1, 95, 10, 99, 2, 9, 99, 103, 1, 5, 103, 107, 1, 5, 107, 111, 2, 111, 10, 115, 1, 6, 115, 119, 2, 10, 119, 123, 1, 6, 123, 127, 1, 127, 5, 131, 2, 9, 131, 135, 1, 5, 135, 139, 1, 139, 10, 143, 1, 143, 2, 147, 1, 147, 5, 0, 99, 2, 0, 14, 0}
		intcode[1] = 12
		intcode[2] = 2
		got := intcode.process()
		want := 5305097
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}

	})

	t.Run("test findNoundAndVerb", func(t *testing.T) {
		intcode := Intcode{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 13, 19, 1, 9, 19, 23, 1, 6, 23, 27, 2, 27, 9, 31, 2, 6, 31, 35, 1, 5, 35, 39, 1, 10, 39, 43, 1, 43, 13, 47, 1, 47, 9, 51, 1, 51, 9, 55, 1, 55, 9, 59, 2, 9, 59, 63, 2, 9, 63, 67, 1, 5, 67, 71, 2, 13, 71, 75, 1, 6, 75, 79, 1, 10, 79, 83, 2, 6, 83, 87, 1, 87, 5, 91, 1, 91, 9, 95, 1, 95, 10, 99, 2, 9, 99, 103, 1, 5, 103, 107, 1, 5, 107, 111, 2, 111, 10, 115, 1, 6, 115, 119, 2, 10, 119, 123, 1, 6, 123, 127, 1, 127, 5, 131, 2, 9, 131, 135, 1, 5, 135, 139, 1, 139, 10, 143, 1, 143, 2, 147, 1, 147, 5, 0, 99, 2, 0, 14, 0}
		testCases := []struct {
			input    int
			expected NounAndVerb
		}{
			{19690720, NounAndVerb{49, 25}},
		}

		for _, tt := range testCases {
			result := intcode.findNounAndVerb(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v want %v", tt.input, tt.expected)
			}
		}
	})
}
