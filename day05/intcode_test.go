package main

import "testing"

func TestIntcoder(t *testing.T) {
	t.Run("test process", func(t *testing.T) {
		testCases := []struct {
			input, expected Intcode
		}{
			{Intcode{memory: []int{1, 0, 0, 0, 99}}, Intcode{memory: []int{2, 0, 0, 0, 99}}},          // (1 + 1 = 2)
			{Intcode{memory: []int{2, 3, 0, 3, 99}}, Intcode{memory: []int{2, 3, 0, 6, 99}}},          // (3 * 2 = 6)
			{Intcode{memory: []int{2, 4, 4, 5, 99, 0}}, Intcode{memory: []int{2, 4, 4, 5, 99, 9801}}}, // (99 * 99 = 9801)
			{Intcode{memory: []int{1, 1, 1, 4, 99, 5, 6, 0, 99}}, Intcode{memory: []int{30, 1, 1, 4, 2, 5, 6, 0, 99}}},
			{Intcode{memory: []int{1101, 100, -1, 4, 0}}, Intcode{memory: []int{1101, 100, -1, 4, 99}}},
		}
		for _, tt := range testCases {
			tt.input.Process()
			equals(t, tt.expected, tt.input)
		}
	})
}
