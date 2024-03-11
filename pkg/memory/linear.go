package memory

import (
	"github.com/d1360-64rc14/risc-v-emulator/pkg/interfaces"
	"github.com/d1360-64rc14/risc-v-emulator/pkg/types"
)

// Linear implements interfaces.Memory
var _ interfaces.Memory[uint32] = (*Linear[uint32])(nil)

// Linear memory space.
type Linear[Arch types.Architecture] struct {
	data []byte
	size Arch
}

func NewLinear[Arch types.Architecture](size Arch) *Linear[Arch] {
	return &Linear[Arch]{
		data: make([]byte, size),
		size: size,
	}
}

func (m *Linear[Arch]) Load(addr Arch) byte {
	return m.data[addr%m.size]
}

func (m *Linear[Arch]) Store(addr Arch, data byte) {
	m.data[addr%m.size] = data
}
