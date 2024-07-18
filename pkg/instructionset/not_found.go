package instructionset

import (
	"fmt"

	"github.com/d1360-64rc14/risc-v-emulator/pkg/memory"
	"github.com/d1360-64rc14/risc-v-emulator/pkg/registerset"
)

// NotFound identifies an undefined instruction.
func NotFound(_ registerset.Register[uint32], pc *uint32, _ memory.Memory[uint32], inst uint32) {
	fmt.Printf("at %X: not found inst %b\n", *pc, inst)
}
