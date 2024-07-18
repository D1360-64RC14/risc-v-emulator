package main

import (
	"fmt"

	"github.com/d1360-64rc14/risc-v-emulator/pkg/base/rv32i"
	"github.com/d1360-64rc14/risc-v-emulator/pkg/memory/linear32bitsmem"
	"github.com/d1360-64rc14/risc-v-emulator/pkg/registerset/base32bits"
)

func main() {
	fmt.Println("Hello RISC-V!")

	x32Regs32bits := base32bits.New()

	linear512ByteMem := linear32bitsmem.New(1024)
	linear512ByteMem.StoreAll(
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
