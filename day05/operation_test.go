package main

import "testing"

func TestOperation(t *testing.T) {
	t.Run("build normal operation", func(t *testing.T) {
		intCode := Intcode{memory: []int{2, 4, 4, 0, 6, 99}}
		o := NewOperation(0, &intCode)
		o.Execute(&intCode)
		equals(t, o.param1, 6)
		equals(t, o.param2, 6)
		equals(t, []int{36, 4, 4, 0, 6, 99}, intCode.memory)
	})

	t.Run("build parameter mode operation", func(t *testing.T) {
		intCode := Intcode{memory: []int{1002, 2, 3, 3, 99}}
		o := NewOperation(0, &intCode)
		o.Execute(&intCode)
		equals(t, o.param1, 3)
		equals(t, o.param2, 3)
		equals(t, o.destination, 3)
		equals(t, intCode.memory, []int{1002, 2, 3, 9, 99})
	})
}

func TestParsingMode(t *testing.T) {
	testCases := []struct {
		input int
		want  map[int]string
	}{
		{1002, map[int]string{
			0: "2",
			1: "position",
			2: "immediate",
			3: "position"}},
		{11002, map[int]string{
			0: "2",
			1: "position",
			2: "immediate",
			3: "immediate"}},
		{10001, map[int]string{
			0: "1",
			1: "position",
			2: "position",
			3: "immediate"}},
		{101, map[int]string{
			0: "1",
			1: "immediate",
			2: "position",
			3: "position"}},
		{104, map[int]string{
			0: "4",
			1: "immediate",
			2: "position",
			3: "position"}},
		{04, map[int]string{
			0: "4",
			1: "position",
			2: "position",
			3: "position"}},
		{4, map[int]string{
			0: "4",
			1: "position",
			2: "position",
			3: "position"}},
	}
	for _, tc := range testCases {
		got := parseOpCode(tc.input)
		equals(t, got, tc.want)
	}
}
