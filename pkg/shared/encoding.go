package shared

// Type mapping
const (
	R_TYPE = 0b111111_000000_00000_111_00000_1111111 // Register
	I_TYPE = 0b111_00000_1111111                     // Immediate
	S_TYPE = 0b111_00000_1111111                     // Store
	B_TYPE = 0b111_00000_1111111                     // Branch
	U_TYPE = 0b1111111                               // Upper Immediate
	J_TYPE = 0b1111111                               // Jump
)
