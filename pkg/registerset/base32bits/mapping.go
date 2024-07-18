package base32bits

// Register mapping for base 32 bits register set
const (
	X0, ZERO byte = iota, iota // Hard-wired zero
	X1, RA   byte = iota, iota // Return address
	X2, SP   byte = iota, iota // Stack pointer
	X3, GP   byte = iota, iota // Global pointer
	X4, TP   byte = iota, iota // Thread pointer
	X5, T0   byte = iota, iota // Temporary
	X6, T1   byte = iota, iota // Temporary
	X7, T2   byte = iota, iota // Temporary
	X8, S0   byte = iota, iota // Saved register/frame pointer
	X9, S1   byte = iota, iota // Saved register
	X10, A0  byte = iota, iota // Function argument/return value
	X11, A1  byte = iota, iota // Function argument/return value
	X12, A2  byte = iota, iota // Function argument
	X13, A3  byte = iota, iota // Function argument
	X14, A4  byte = iota, iota // Function argument
	X15, A5  byte = iota, iota // Function argument
	X16, A6  byte = iota, iota // Function argument
	X17, A7  byte = iota, iota // Function argument
	X18, S2  byte = iota, iota // Saved register
	X19, S3  byte = iota, iota // Saved register
	X20, S4  byte = iota, iota // Saved register
	X21, S5  byte = iota, iota // Saved register
	X22, S6  byte = iota, iota // Saved register
	X23, S7  byte = iota, iota // Saved register
	X24, S8  byte = iota, iota // Saved register
	X25, S9  byte = iota, iota // Saved register
	X26, S10 byte = iota, iota // Saved register
	X27, S11 byte = iota, iota // Saved register
	X28, T3  byte = iota, iota // Temporary
	X29, T4  byte = iota, iota // Temporary
	X30, T5  byte = iota, iota // Temporary
	X31, T6  byte = iota, iota // Temporary
)
