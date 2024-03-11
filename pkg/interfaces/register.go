package interfaces

import "github.com/d1360-64rc14/risc-v-emulator/pkg/types"

type RegisterMapping interface {
	types.X32Regs
}

type Register[Arch types.Architecture, RegMap RegisterMapping] interface {
	Get(reg RegMap) Arch
	Set(reg RegMap, value Arch)
}
