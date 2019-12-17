package main

func main() {
	var memory []int = parseInputs("input.txt")
	intcode := Intcode{memory: memory}
	intcode.input = getUserInput()
	intcode.Process()
}
