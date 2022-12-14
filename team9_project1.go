package main

import (
	"flag"
)

type Instruction struct {
	rawInstruction  string
	lineValue       uint64
	memLoc          uint64
	opcode          uint64
	op              string
	instructionType string
	rm              uint8
	shamt           uint8
	rn              uint8
	rd              uint8
	rt              uint8
	op2             uint8
	address         uint16
	immediate       int16
	offset          int32
	conditional     uint8
	shiftCode       uint8
	field           uint32
	memValue        int64
}

var InstructionList []Instruction

func main() {
	inputFilePathPtr := flag.String("i", "addtest1_bin.txt", "input file path")
	outputFilePathPtr := flag.String("o", "out.txt", "output file path")

	flag.Parse()
	//Inputs Command-Line
	ReadBinary(*inputFilePathPtr)

	ProcessInstructionList(InstructionList)
	//fmt.Println(InstructionList)
	WriteInstructions(*outputFilePathPtr, InstructionList)

}
