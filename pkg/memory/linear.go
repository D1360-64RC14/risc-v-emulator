package memory

import (
	"github.com/d1360-64rc14/risc-v-emulator/pkg/interfaces"
	"golang.org/x/exp/constraints"
)

// Linear implements interfaces.Memory
var _ interfaces.Memory[uint32] = (*Linear[uint32])(nil)

// Linear memory space.
type Linear[A constraints.Unsigned] struct {
	data []byte
	size A
}

func NewLinear[A constraints.Unsigned](size A) *Linear[A] {
	return &Linear[A]{
		data: make([]byte, size),
		size: size,
	}
}

func (c *Linear[A]) Load(addr A) byte {
	return c.data[addr%c.size]
}

func (c *Linear[A]) Store(addr A, data byte) {
	c.data[addr%c.size] = data
}
