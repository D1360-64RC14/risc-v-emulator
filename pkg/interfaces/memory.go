package interfaces

import "github.com/d1360-64rc14/risc-v-emulator/pkg/types"

type Memory[Arch types.Architecture] interface {
	Load(addr Arch) byte
	Store(addr Arch, data byte)
}
