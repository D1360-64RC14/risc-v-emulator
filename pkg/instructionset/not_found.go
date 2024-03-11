package instructionset

import (
	"fmt"

	"github.com/d1360-64rc14/risc-v-emulator/pkg/interfaces"
	"github.com/d1360-64rc14/risc-v-emulator/pkg/types"
)

// NotFound identifies an undefined instruction.
func NotFound(_ interfaces.Register[uint32, types.X32Regs], pc *uint32, _ interfaces.Memory[uint32], inst uint32) {
	fmt.Printf("at %X: not found inst %b\n", *pc, inst)
}
