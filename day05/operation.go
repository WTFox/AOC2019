package main

import (
	"fmt"
	"log"
	"strconv"
)

type Operation struct {
	code,
	param1,
	param2,
	destination,
	size int
}

func (o Operation) Execute(i *Intcode) {
	switch o.code {
	case 1:
		i.memory[o.destination] = o.param1 + o.param2
	case 2:
		i.memory[o.destination] = o.param1 * o.param2
	case 3:
		i.memory[o.destination] = i.input
	case 4:
		i.output = i.memory[o.destination]
		fmt.Println(i.output)
	}
}

func parseOpCode(opcode int) map[int]string {
	result := make(map[int]string)
	if opcode < 10 {
		return map[int]string{
			0: strconv.Itoa(opcode),
			1: "position",
			2: "position",
			3: "position",
		}
	}
	opcodeSlice := convertNumberToSlice(opcode)

	result[0] = strconv.Itoa(opcodeSlice[len(opcodeSlice)-1])

	for idx := range opcodeSlice[:len(opcodeSlice)-2] {
		reverseIndex := len(opcodeSlice) - idx - 3
		val := opcodeSlice[reverseIndex]
		switch val {
		case 1:
			result[idx+1] = "immediate"
		default:
			result[idx+1] = "position"
		}
	}
	if len(result) == 2 {
		result[2] = "position"
		result[3] = "position"
	}
	if len(result) == 3 {
		result[3] = "position"
	}
	return result
}

func buildParametrizedOperation(i *Intcode, metadata map[int]string, address, size int) Operation {
	instructions := i.memory[address : address+size]
	code, err := strconv.Atoi(metadata[0])
	if err != nil {
		log.Fatalf("Couldn't parse instruction: %v\n", metadata[0])
	}

	var param1 int
	if metadata[1] == "position" {
		param1 = i.memory[instructions[1]]
	} else {
		param1 = instructions[1]
	}

	if size == 2 {
		return Operation{
			size:        size,
			code:        code,
			destination: instructions[1],
		}
	}

	var param2 int
	if metadata[2] == "position" {
		param2 = i.memory[instructions[2]]
	} else {
		param2 = instructions[2]
	}

	return Operation{
		size:        size,
		code:        code,
		param1:      param1,
		param2:      param2,
		destination: instructions[3],
	}

}

func buildShortOperation(i *Intcode, metadata map[int]string, address, size int) Operation {
	instructions := i.memory[address : address+size]
	code, err := strconv.Atoi(metadata[0])
	if err != nil {
		log.Fatalf("Couldn't parse instruction: %v\n", metadata[0])
	}

	return Operation{
		size:        size,
		code:        code,
		destination: instructions[3],
	}
}

func NewOperation(address int, i *Intcode) Operation {
	opcodeMetadata := parseOpCode(i.memory[address])
	op, _ := strconv.Atoi(opcodeMetadata[0])

	var size int
	switch op {
	case 3, 4:
		size = 2
	default:
		size = 4
	}
	return buildParametrizedOperation(i, opcodeMetadata, address, size)
}
