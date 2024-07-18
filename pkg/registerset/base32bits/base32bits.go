package base32bits

import (
	"fmt"

	"github.com/d1360-64rc14/risc-v-emulator/pkg/registerset"
)

// [Base32Bits] implements [registerset.Register]
var _ registerset.Register[uint32] = (*Base32Bits)(nil)

type Base32Bits [32]uint32

func New() *Base32Bits {
	return &Base32Bits{}
}

func (x Base32Bits) Get(reg byte) uint32 {
	if reg == X0 {
		return 0
	}

	return x[reg]
}

func (x *Base32Bits) Set(reg byte, value uint32) {
	if reg == X0 {
		return
	}

	x[reg] = value
}

func (x Base32Bits) Dump() {
	fmt.Println("=== x32register.X32Register - Dump ===")
	for i := 0; i < 32; i++ {
		fmt.Printf("%2d: %08x - %d\n", i, x[i], x[i])
	}
}
