package instruction

import "github.com/d1360-64rc14/risc-v-emulator/internal/memory"

type Instruction func(regs [32]uint32, pc *uint32, mem *memory.Memory, inst uint32)
