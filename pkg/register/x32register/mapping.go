package x32register

import (
	"github.com/d1360-64rc14/risc-v-emulator/pkg/types"
)

// Register mapping
const (
	X0, ZERO types.X32Regs = iota, iota // Hard-wired zero
	X1, RA   types.X32Regs = iota, iota // Return address
	X2, SP   types.X32Regs = iota, iota // Stack pointer
	X3, GP   types.X32Regs = iota, iota // Global pointer
	X4, TP   types.X32Regs = iota, iota // Thread pointer
	X5, T0   types.X32Regs = iota, iota // Temporary
	X6, T1   types.X32Regs = iota, iota // Temporary
	X7, T2   types.X32Regs = iota, iota // Temporary
	X8, S0   types.X32Regs = iota, iota // Saved register/frame pointer
	X9, S1   types.X32Regs = iota, iota // Saved register
	X10, A0  types.X32Regs = iota, iota // Function argument/return value
	X11, A1  types.X32Regs = iota, iota // Function argument/return value
	X12, A2  types.X32Regs = iota, iota // Function argument
	X13, A3  types.X32Regs = iota, iota // Function argument
	X14, A4  types.X32Regs = iota, iota // Function argument
	X15, A5  types.X32Regs = iota, iota // Function argument
	X16, A6  types.X32Regs = iota, iota // Function argument
	X17, A7  types.X32Regs = iota, iota // Function argument
	X18, S2  types.X32Regs = iota, iota // Saved register
	X19, S3  types.X32Regs = iota, iota // Saved register
	X20, S4  types.X32Regs = iota, iota // Saved register
	X21, S5  types.X32Regs = iota, iota // Saved register
	X22, S6  types.X32Regs = iota, iota // Saved register
	X23, S7  types.X32Regs = iota, iota // Saved register
	X24, S8  types.X32Regs = iota, iota // Saved register
	X25, S9  types.X32Regs = iota, iota // Saved register
	X26, S10 types.X32Regs = iota, iota // Saved register
	X27, S11 types.X32Regs = iota, iota // Saved register
	X28, T3  types.X32Regs = iota, iota // Temporary
	X29, T4  types.X32Regs = iota, iota // Temporary
	X30, T5  types.X32Regs = iota, iota // Temporary
	X31, T6  types.X32Regs = iota, iota // Temporary
)
