package instruction

import (
	"fmt"

	"github.com/d1360-64rc14/risc-v-emulator/internal/memory"
)

func NotFound(_ [32]uint32, pc *uint32, mem *memory.Memory, inst uint32) {
	fmt.Printf("at %X: not found inst %b\n", *pc, inst)
}
