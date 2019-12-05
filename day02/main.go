package main

// Intcode ...
type Intcode []int

// NounAndVerb ...
type NounAndVerb struct{ noun, verb int }

func (i Intcode) findNounAndVerb(want int) NounAndVerb {
	for a := 0; a < 100; a++ {
		for b := 0; b < 100; b++ {
			intcodeCopy := make(Intcode, len(i))
			copy(intcodeCopy, i)
			intcodeCopy[1] = a
			intcodeCopy[2] = b
			if result := intcodeCopy.process(); result == want {
				return NounAndVerb{intcodeCopy[1], intcodeCopy[2]}
			}
		}
	}
	return NounAndVerb{-1, -1}
}

func (i *Intcode) process() int {
	// golang doesn't support indexing on referenced objects
	// but I need the pointer reference to update values
	// in place. :thisisfine:
	self := *i

	address := 0
	for {
		var opcode int
		if opcode = self[address]; opcode == 99 {
			break
		}

		parameterOne := self[self[address+1]]
		parameterTwo := self[self[address+2]]
		outputDestination := self[address+3]

		if outputDestination > len(self) {
			return -1
		}

		switch opcode {
		case 1:
			self[outputDestination] = parameterOne + parameterTwo
		case 2:
			self[outputDestination] = parameterOne * parameterTwo
		}

		// moving along
		address += 4
	}

	// output is stored at address 0
	return self[0]
}

func main() {
	return
}
