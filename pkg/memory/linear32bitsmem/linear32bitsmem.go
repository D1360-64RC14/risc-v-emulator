package linear32bitsmem

import (
	"fmt"

	"github.com/d1360-64rc14/risc-v-emulator/pkg/memory"
)

// [Linear32BitsMemory] implements [memory.Memory]
var _ memory.Memory[uint32] = (*Linear32BitsMemory)(nil)

// [Linear32BitsMemory] memory space.
type Linear32BitsMemory struct {
	data []byte
	size uint32
}

func New(size uint32) *Linear32BitsMemory {
	return &Linear32BitsMemory{
		data: make([]byte, size),
		size: size,
	}
}

func (m Linear32BitsMemory) Load(addr uint32) byte {
	return m.data[addr%m.size]
}

func (m *Linear32BitsMemory) Store(addr uint32, data byte) {
	m.data[addr%m.size] = data
}

func (m *Linear32BitsMemory) StoreAll(data ...uint32) {
	for i := 0; i < len(data); i++ {
		m.Store(uint32(i*4), byte(data[i]))
		m.Store(uint32(i*4+1), byte(data[i]>>8))
		m.Store(uint32(i*4+2), byte(data[i]>>16))
		m.Store(uint32(i*4+3), byte(data[i]>>24))
	}
}

func (m *Linear32BitsMemory) Dump() {
	fmt.Println("=== memory.Linear - Dump ===")
	for i := 0; i < len(m.data); i += 4 {
		fmt.Printf("%08x: %08x %08x %08x %08x\n", i, m.data[i], m.data[i+1], m.data[i+2], m.data[i+3])
	}
}
