package memory

import (
	"golang.org/x/exp/constraints"
)

// Contiguous memory space.
type Contignuous[A constraints.Unsigned] /* implements interfaces.Memory */ struct {
	data []byte
	size A
}

func NewContignuous[A constraints.Unsigned](size A) *Contignuous[A] {
	return &Contignuous[A]{
		data: make([]byte, size),
		size: size,
	}
}

func (c *Contignuous[A]) Load(addr A) byte {
	return c.data[addr%c.size]
}

func (c *Contignuous[A]) Store(addr A, data byte) {
	c.data[addr%c.size] = data
}
