package main

import (
	"fmt"
	"sort"
)

var testInputLow = 254032
var testInputHigh = 789860

// 1033

/*
However, they do remember a few key facts about the password:

    - It is a six-digit number. (duh)
	- The value is within the range given in your puzzle input. (duh)

    - Two adjacent digits are the same (like 22 in 122345).
    - Going from left to right, the digits never decrease; they only ever increase or stay the same (like 111123 or 135679).

Other than the range rule, the following are true:

    - 111111 meets these criteria (double 11, never decreases).
    - 223450 does not meet these criteria (decreasing pair of digits 50).
	- 123789 does not meet these criteria (no double).

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
	result := false
	for idx, val := range nums {
		if idx == len(nums)-1 {
			return result
		}
		if val == nums[idx+1] {
			result = true
			break
		}
	}
	return result
}

func getCombinations(low, high int) (combinations int) {
	for i := testInputLow; i <= testInputHigh; i++ {
		nums := convertNumberToSlice(int(i))
		if !sort.IntsAreSorted(nums) {
			continue
		}
		if !hasAdjacentPair(nums) {
			continue
		}
		combinations++
	}
	return
}

func main() {
	fmt.Println(getCombinations(testInputLow, testInputHigh))
}
