package instructionset

import (
	"fmt"

	"github.com/d1360-64rc14/risc-v-emulator/pkg/interfaces"
)

// NotFound identifies an undefined instruction.
func NotFound(_ *[32]uint32, pc *uint32, _ interfaces.Memory[uint32], inst uint32) {
	fmt.Printf("at %X: not found inst %b\n", *pc, inst)
}
