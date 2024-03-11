package main

import (
	"fmt"

	"github.com/d1360-64rc14/risc-v-emulator/pkg/base/rv32i"
	"github.com/d1360-64rc14/risc-v-emulator/pkg/memory"
	"github.com/d1360-64rc14/risc-v-emulator/pkg/register/x32register"
)

func main() {
	fmt.Println("Hello RISC-V!")

	x32Regs32bits := x32register.New[uint32]()

	linear512ByteMem := memory.NewLinear[uint32](1024)
	linear512ByteMem.StoreAllU32(
		0x00a00293,
		0x01400313,
		0x01e00393,
		0x00100073,
	)

	baseRv32i := rv32i.New(x32Regs32bits, linear512ByteMem)

	baseRv32i.Start()

	x32Regs32bits.Dump()
	linear512ByteMem.Dump()
}
