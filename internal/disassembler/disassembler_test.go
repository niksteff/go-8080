package disassembler_test

import (
	"os"
	"testing"

	"github.com/niksteff/go-8080/internal/disassembler"
)

type FuncWriter func(b []byte) (n int, err error)

func (f FuncWriter) Write(b []byte) (n int, err error) {
	return f(b)
}

func TestDisassemblerToTarget(t *testing.T) {
	// load the test file
	f, err := os.Open("TST8080.COM")
	if err != nil {
		t.Fatalf("error opening test file: %v", err)
	}
	defer f.Close()

	w := FuncWriter(func(b []byte) (int, error) {
		t.Logf(string(b))
		
		return len(b), nil
	})

	disassembler.ReadProgram(f, w)
}

func TestDisassemblerToFile(t *testing.T) {
	// load the test file
	f, err := os.Open("TST8080.COM")
	if err != nil {
		t.Fatalf("error opening test file: %v", err)
	}
	defer f.Close()

	tf, err := os.Create("TST8080.asm")
	if err != nil {
		t.Fatalf("error creating output file: %v", err)
	}
	defer tf.Close()

	w := FuncWriter(func(b []byte) (int, error) {
		_, err = tf.Write(b)
		if err != nil {
			t.Fatalf("error writing to output file: %v", err)
		}

		return len(b), nil
	})

	disassembler.ReadProgram(f, w)
}
