package x32register

import (
	"fmt"

	"github.com/d1360-64rc14/risc-v-emulator/pkg/interfaces"
	"github.com/d1360-64rc14/risc-v-emulator/pkg/types"
)

// X32Register implements interfaces.Register
var _ interfaces.Register[uint32, types.X32Regs] = (*X32Register[uint32])(nil)

type X32Register[Arch types.Architecture] struct {
	regs [32]Arch
}

func New[Arch types.Architecture]() *X32Register[Arch] {
	return &X32Register[Arch]{}
}

func (x *X32Register[Arch]) Get(reg types.X32Regs) Arch {
	if reg == X0 {
		return 0
	}

	return x.regs[reg]
}

func (x *X32Register[Arch]) Set(reg types.X32Regs, value Arch) {
	if reg == X0 {
		return
	}

	x.regs[reg] = value
}

func (x *X32Register[Arch]) Dump() {
	fmt.Println("=== x32register.X32Register - Dump ===")
	for i := 0; i < 32; i++ {
		fmt.Printf("%2d: %08x - %d\n", i, x.regs[i], x.regs[i])
	}
}
