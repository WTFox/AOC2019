package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseInputs(filename string) []int {
	inputBytes, _ := ioutil.ReadFile(filename)
	testInput := strings.Split(string(inputBytes), "\n")[0]
	testInputs := strings.Split(testInput, ",")

	output := make([]int, 0)

	for _, val := range testInputs {
		if val, err := strconv.Atoi(string(val)); err == nil {
			output = append(output, val)
		} else {
			log.Fatalf("Couldn't parse input.txt\n%v\n", err)
		}
	}
	return output
}

func convertNumberToSlice(num int) []int {
	output := make([]int, 0)
	for num > 0 {
		var digit int
		digit = num % 10
		output = append([]int{digit}, output...)
		num = num / 10
	}
	return output
}

func getUserInput() int {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	number, _ := strconv.Atoi(text)
	return number
}
