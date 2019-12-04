package main

import (
	"reflect"
	"testing"
)

func TestIntcoder(t *testing.T) {
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
}
