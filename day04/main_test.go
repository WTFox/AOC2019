package main

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"testing"
)

func TestConvertNumberToArray(t *testing.T) {
	testCases := []struct {
		input int
		want  []int
	}{
		{254032, []int{2, 5, 4, 0, 3, 2}},
		{8675309, []int{8, 6, 7, 5, 3, 0, 9}},
		{867530, []int{8, 6, 7, 5, 3, 0}},
		{217530, []int{2, 1, 7, 5, 3, 0}},
	}

	for _, tc := range testCases {
		got := convertNumberToSlice(tc.input)
		equals(t, tc.want, got)
	}
}

func TestAdjacentPair(t *testing.T) {
	testCases := []struct {
		input []int
		want  bool
	}{
		{[]int{2, 5, 4, 0, 3, 2}, false},
		{[]int{8, 6, 7, 5, 3, 0, 9}, false},
		{[]int{8, 6, 1, 1, 3, 0, 9}, true},
		{[]int{8, 6, 3, 1, 1}, true},
		{[]int{1, 1, 1, 1, 1}, false},
		{[]int{1, 1, 1, 6, 8}, false},
		{[]int{8, 6, 1, 1, 1, 0}, false},
		{[]int{1, 2, 3, 4, 4, 4}, false},
		{[]int{1, 1, 1, 1, 2, 2}, true},
		{[]int{1, 1, 2, 2, 3, 3}, true},
		{[]int{7, 8, 8, 9, 9, 9}, true},
	}
	for _, tc := range testCases {
		fmt.Printf("Testing %v\n", tc.input)
		got := hasAdjacentPair(tc.input)
		equals(t, tc.want, got)
	}
}

func TestIntsSorted(t *testing.T) {
	testCases := []struct {
		input []int
		want  bool
	}{
		{[]int{1, 2, 3, 4, 5, 0}, false},
		{[]int{1, 2, 3}, true},
		{[]int{3, 2, 1}, false},
	}
	for _, tc := range testCases {
		got := sort.IntsAreSorted(tc.input)
		equals(t, tc.want, got)
	}
}

func TestValidateNumber(t *testing.T) {
	testCases := []struct {
		input int
		want  bool
	}{
		{223450, false},
		{123789, false},
		{11233, true},
		{123444, false},
		{111122, true},
	}

	for _, tc := range testCases {
		got := validateNumber(tc.input)
		equals(t, tc.want, got)
	}
}

func TestGetCombinations(t *testing.T) {
	testCases := []struct {
		inputLow  int
		inputHigh int
		want      int
	}{
		{11121, 11135, 2}, // 22, 33
		{11121, 11135, 2}, // 22, 33
		{11121, 11135, 2}, // 22, 33
	}

	for _, tc := range testCases {
		got := getCombinations(tc.inputLow, tc.inputHigh)
		equals(t, tc.want, got)
	}
}

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
