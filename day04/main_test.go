package main

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestConvertNumberToArray(t *testing.T) {
	testCases := []struct {
		input int
		want  []int
	}{
		{254032, []int{2, 5, 4, 0, 3, 2}},
		{8675309, []int{8, 6, 7, 5, 3, 0, 9}},
	}

	for _, tc := range testCases {
		got := convertNumberToSlice(tc.input)
		equals(t, got, tc.want)
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
	}
	for _, tc := range testCases {
		got := hasAdjacentPair(tc.input)
		equals(t, got, tc.want)
	}
}

func TestGetCombinations(t *testing.T) {
	got := getCombinations(testInputLow, testInputHigh)
	want := 1033
	equals(t, got, want)
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
