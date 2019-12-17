package main

type Intcode struct {
	input,
	output int
	memory []int
}

func (i *Intcode) Process() {
	address := 0
	for {
		if opcode := i.memory[address]; opcode == 99 {
			break
		}
		op := NewOperation(address, i)
		op.Execute(i)
		address += op.size
	}
	return
}
