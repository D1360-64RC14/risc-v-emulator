package riscv

// Register mapping
const (
	X0, ZERO = iota, iota // Hard-wired zero
	X1, RA   = iota, iota // Return address
	X2, SP   = iota, iota // Stack pointer
	X3, GP   = iota, iota // Global pointer
	X4, TP   = iota, iota // Thread pointer
	X5, T0   = iota, iota // Temporary
	X6, T1   = iota, iota // Temporary
	X7, T2   = iota, iota // Temporary
	X8, S0   = iota, iota // Saved register/frame pointer
	X9, S1   = iota, iota // Saved register
	X10, A0  = iota, iota // Function argument/return value
	X11, A1  = iota, iota // Function argument/return value
	X12, A2  = iota, iota // Function argument
	X13, A3  = iota, iota // Function argument
	X14, A4  = iota, iota // Function argument
	X15, A5  = iota, iota // Function argument
	X16, A6  = iota, iota // Function argument
	X17, A7  = iota, iota // Function argument
	X18, S2  = iota, iota // Saved register
	X19, S3  = iota, iota // Saved register
	X20, S4  = iota, iota // Saved register
	X21, S5  = iota, iota // Saved register
	X22, S6  = iota, iota // Saved register
	X23, S7  = iota, iota // Saved register
	X24, S8  = iota, iota // Saved register
	X25, S9  = iota, iota // Saved register
	X26, S10 = iota, iota // Saved register
	X27, S11 = iota, iota // Saved register
	X28, T3  = iota, iota // Temporary
	X29, T4  = iota, iota // Temporary
	X30, T5  = iota, iota // Temporary
	X31, T6  = iota, iota // Temporary
)
