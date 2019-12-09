package main

import (
	"fmt"
	"math"
	"sort"
)

var testInputLow = 254032
var testInputHigh = 789860

/*
However, they do remember a few key facts about the password:

    - It is a six-digit number. (duh)
	- The value is within the range given in your puzzle input. (duh)

    - Two adjacent digits are the same (like 22 in 122345).
    - Going from left to right, the digits never decrease; they only ever increase or stay the same (like 111123 or 135679).

Other than the range rule, the following are true:

    - 223450 does not meet these criteria (decreasing pair of digits 50).
	- 123789 does not meet these criteria (no double).
    - 112233 meets these criteria because the digits never decrease and all repeated digits are exactly two digits long.
    - 123444 doesn't meet the criteria (the repeated 44 is part of a larger group of 444).
    - 111122 meets the criteria (even though 1 is repeated more than twice, it still contains a double 22).


How many different passwords within the range given in your puzzle input meet these criteria?
*/

func convertNumberToSlice(num int) []int {
	output := make([]int, 0)
	for num > 0 {
		digit := num % 10
		output = append([]int{digit}, output...)
		num = num / 10
	}
	return output
}

func hasAdjacentPair(nums []int) bool {
	mapCount := make(map[int][]int) // val:count
	for idx, val := range nums {
		mapCount[val] = append(mapCount[val], idx)
	}
	for _, v := range mapCount {
		if len(v) == 2 {
			left := float64(v[0])
			right := float64(v[1])
			if math.Max(left, right)-math.Min(left, right) == 1 {
				return true
			}
		}
	}
	return false
}

func validateNumber(num int) bool {
	nums := convertNumberToSlice(num)
	if !hasAdjacentPair(nums) {
		return false
	}
	if !sort.IntsAreSorted(nums) {
		return false
	}
	return true
}

func getCombinations(low, high int) int {
	combinations := 0
	for i := low; i <= high; i++ {
		if validateNumber(i) {
			combinations++
		}
	}
	return combinations
}

func main() {
	fmt.Println(getCombinations(testInputLow, testInputHigh))
}
