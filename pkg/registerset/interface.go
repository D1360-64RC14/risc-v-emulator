package registerset

import "github.com/d1360-64rc14/risc-v-emulator/pkg/types"

type Register[Arch types.Architecture] interface {
	Get(reg byte) Arch
	Set(reg byte, value Arch)
}
