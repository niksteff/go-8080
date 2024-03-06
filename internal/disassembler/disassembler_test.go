package disassembler_test

import (
	"os"
	"testing"

	"github.com/niksteff/go-8080/internal/disassembler"
)

func TestDisassembler(t *testing.T) {
	// load the test file
	f, err := os.Open("TST8080.COM")
	if err != nil {
		t.Fatalf("error opening test file: %v", err)
	}
	defer f.Close()

	disassembler.ReadProgram(f)
}