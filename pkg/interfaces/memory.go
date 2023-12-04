package interfaces

import "golang.org/x/exp/constraints"

type Memory[A constraints.Unsigned] interface {
	Load(addr A) byte
	Store(addr A, data byte)
}
