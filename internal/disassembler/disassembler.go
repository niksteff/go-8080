package disassembler

import (
	"fmt"
	"io"
	"log"
)

// TODO: can i rewrite this using a buffer and a peek advance pattern?
// - read a byte
// - check which instruction it is
// - peek the remaining byte instructions
// - advance the buffer by the number of bytes read
// - if possible do not read the entire file into memory

func ReadProgram(r io.Reader) {
	log.Printf("disassembling program...\n ")

	// first read in the entire program
	buf, err := io.ReadAll(r)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	log.Printf("read %d bytes\n\n", len(buf))

	var pc uint = 0 // our program counter starts at 0

	// iterate through the program and disassemble each instruction
	// the pc may jump multiple bytes depending how many bytes the instruction
	// read.
	for pc < uint(len(buf)) {
		pc += Disassemble(buf, pc)
	}

	log.Printf("\n disassembled %d bytes\n", pc)
}

func Disassemble(b []byte, pc uint) uint {
	var ins instruction
	ins.opcode = b[pc]

	switch b[pc] {
	case 0x00:
		ins.name = "NOP"
	case 0xc3:
		ins.name = "JMP"
		ins.parameters = append(ins.parameters, b[pc+2], b[pc+1])
	default:
		ins.name = "UNKWN"
	}

	log.Printf("%04x	\t%s", pc, ins.String())

	return uint(len(ins.parameters) + 1)
}

type instruction struct {
	name       string
	opcode     byte   // the OP code for the instruction
	parameters []byte // how many parameters the instruction takes
}

func (i *instruction) String() string {
	s := fmt.Sprintf("%02x ", i.opcode)

	for _, p := range i.parameters {
		s += fmt.Sprintf("%02x ", p)
	}

	if len(i.parameters) == 0 {
		s += "\t\t"
	}

	switch len(i.parameters) {
	case 1:
		s += fmt.Sprintf("\t	%s	\t%02x", i.name, i.parameters[0])
	case 2:
		s += fmt.Sprintf("\t	%s	\t$%02x%02x", i.name, i.parameters[0], i.parameters[1])
	default:
		// no parameters, just the op code
		s += fmt.Sprintf("\t	%s", i.name)
	}

	return s
}
