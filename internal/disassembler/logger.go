package disassembler

import (
	"fmt"
	"log"
)

func init() {
	log.SetFlags(0)
    log.SetOutput(new(disassemblyLogWriter))
}

type disassemblyLogWriter struct {}

func (writer disassemblyLogWriter) Write(bytes []byte) (int, error) {
    return fmt.Print(string(bytes))
}