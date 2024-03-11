package memory

import (
	"fmt"

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

func (m *Linear[Arch]) StoreAllU32(data ...uint32) {
	for i := 0; i < len(data); i++ {
		m.Store(Arch(i*4), byte(data[i]))
		m.Store(Arch(i*4+1), byte(data[i]>>8))
		m.Store(Arch(i*4+2), byte(data[i]>>16))
		m.Store(Arch(i*4+3), byte(data[i]>>24))
	}
}

func (m *Linear[Arch]) Dump() {
	fmt.Println("=== memory.Linear - Dump ===")
	for i := 0; i < len(m.data); i += 4 {
		fmt.Printf("%08x: %08x %08x %08x %08x\n", i, m.data[i], m.data[i+1], m.data[i+2], m.data[i+3])
	}
}
