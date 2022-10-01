package main

import (
	"fmt"
	"os"
)

type Instruction struct {
	rawInstruction  string
	lineValue       uint64
	opcode          uint64
	op              string
	instructionType string
	rm              uint8
	shamt           uint8
	rn              uint8
	rd              uint8
	rt              uint8
	op2             uint8
	address         uint8
	immediate       uint8
	offset          uint8
	conditional     uint8
	shiftCode       uint8
	field           uint8
}

var InstructionList []Instruction

func main() {
	//Checks if addtest1_bin.txt exists
	if _, err := os.Stat("addtest1_bin.txt"); err == nil {
		os.Args[0] = "addtest1_bin.txt"
	}

	//Inputs Command-Line
	ReadBinary(os.Args[0])

	fmt.Println(InstructionList)
	ProcessInstructionList(InstructionList)
	//fmt.Println(InstructionList)

	//Prints Array
	for i := 0; i < len(InstructionList); i++ {
		fmt.Println()
	}
}
