package rv32i

import (
	"github.com/d1360-64rc14/risc-v-emulator/pkg/interfaces"
	"github.com/d1360-64rc14/risc-v-emulator/pkg/shared"
	"github.com/d1360-64rc14/risc-v-emulator/pkg/types"
)

// LUI (load upper immediate) is used to build 32-bit constants and uses the U-type format. LUI
// places the U-immediate value in the top 20 bits of the destination register rd, filling in the lowest
// 12 bits with zeros.
//
// lui rd, imm
func LUI(regs interfaces.Register[uint32, types.X32Regs], pc *uint32, mem interfaces.Memory[uint32], inst uint32) {
	regs.Set(shared.RD(inst), shared.ImmU(inst))
}

// AUIPC (add upper immediate to pc) is used to build pc-relative addresses and uses the U-type
// format. AUIPC forms a 32-bit offset from the 20-bit U-immediate, filling in the lowest 12 bits with
// zeros, adds this offset to the address of the AUIPC instruction, then places the result in register
// rd.
//
// auipc rd, imm
func AUIPC(regs interfaces.Register[uint32, types.X32Regs], pc *uint32, mem interfaces.Memory[uint32], inst uint32) {
	*pc += shared.ImmU(inst)

	regs.Set(shared.RD(inst), *pc)
}

// JAL (jump and link) is used to jump to a new immediate location and writes the address of the following
// instruction to rd.
// The address of the instruction following the jump (pc+4) is written to register rd.
//
// jal rd, imm
func JAL(regs interfaces.Register[uint32, types.X32Regs], pc *uint32, mem interfaces.Memory[uint32], inst uint32) {
	*pc += shared.ImmJ(inst)

	regs.Set(shared.RD(inst), *pc+4)
}

// JALR (jump and link reigster) is used to jump to a new register + immediate location and writes the
// address of the following instruction to rd.
// The target address is obtained by adding the sign-extended 12-bit immediate to the register rs1, then setting
// the least-significant bit of the result to zero. The address of the instruction following the jump
// (pc+4) is written to register rd.
//
// jalr rd, offset(rs1)
func JALR(regs interfaces.Register[uint32, types.X32Regs], pc *uint32, mem interfaces.Memory[uint32], inst uint32) {
	*pc += regs.Get(shared.RS1(inst)) + shared.ImmI(inst)

	regs.Set(shared.RD(inst), *pc+4)
}

// BEQ (branch if equal) take the branch if registers rs1 and rs2 are equal.
//
// beq rs1, rs2, offset
func BEQ(regs interfaces.Register[uint32, types.X32Regs], pc *uint32, mem interfaces.Memory[uint32], inst uint32) {
	rs1 := regs.Get(shared.RS1(inst))
	rs2 := regs.Get(shared.RS2(inst))

	if rs1 == rs2 {
		*pc += shared.ImmI(inst)
	}
}

// BNE (branch if not equal) take the branch if registers rs1 and rs2 are unequal.
//
// bne rs1, rs2, offset
func BNE(regs interfaces.Register[uint32, types.X32Regs], pc *uint32, mem interfaces.Memory[uint32], inst uint32) {
	rs1 := regs.Get(shared.RS1(inst))
	rs2 := regs.Get(shared.RS2(inst))

	if rs1 != rs2 {
		*pc += shared.ImmI(inst)
	}
}

// BLT (branch if less than) take the branch if rs1 is less than rs2, using
// signed comparison.
//
// blt rs1, rs2, offset
func BLT(regs interfaces.Register[uint32, types.X32Regs], pc *uint32, mem interfaces.Memory[uint32], inst uint32) {
	rs1 := regs.Get(shared.RS1(inst))
	rs2 := regs.Get(shared.RS2(inst))

	equalSignBit := rs1>>31 == rs2>>31

	// If equals sign bit, can be compared directly; else, need to do arithmetics to correct number.
	if (equalSignBit && rs1 < rs2) || rs1-rs2 < rs1 {
		*pc += shared.ImmI(inst)
	}
}

// BGE (branch if greater than or equal) take the branch if rs1 is greater
// than or equal to rs2, using signed comparison.
//
// bge rs1, rs2, offset
func BGE(regs interfaces.Register[uint32, types.X32Regs], pc *uint32, mem interfaces.Memory[uint32], inst uint32) {
	rs1 := regs.Get(shared.RS1(inst))
	rs2 := regs.Get(shared.RS2(inst))

	signBitEquals := rs1>>31 == rs2>>31

	// If equals sign bit, can be compared directly; else, need to do arithmetics to correct number.
	if (signBitEquals && rs1 >= rs2) || rs1-rs2 >= rs1 {
		*pc += shared.ImmI(inst)
	}
}

// BLTU (branch if less than unsigned) take the branch if rs1 is less than rs2, using
// unsigned comparison.
//
// bltu rs1, rs2, offset
func BLTU(regs interfaces.Register[uint32, types.X32Regs], pc *uint32, mem interfaces.Memory[uint32], inst uint32) {
	rs1 := regs.Get(shared.RS1(inst))
	rs2 := regs.Get(shared.RS2(inst))

	if rs1 < rs2 {
		*pc += shared.ImmI(inst)
	}
}

// BGEU (branch if greater than or equal unsigned) take the branch if rs1 is greater
// than or equal to rs2, using unsigned comparison.
//
// bgeu rs1, rs2, offset
func BGEU(regs interfaces.Register[uint32, types.X32Regs], pc *uint32, mem interfaces.Memory[uint32], inst uint32) {
	rs1 := regs.Get(shared.RS1(inst))
	rs2 := regs.Get(shared.RS2(inst))

	if rs1 >= rs2 {
		*pc += shared.ImmI(inst)
	}
}

// LB (load byte) loads an 8-bit value from memory at address offset + rs1 to rd.
// Sign-extending to 32-bits before storing in rd.
//
// lb rd, offset(rs1)
func LB(regs interfaces.Register[uint32, types.X32Regs], pc *uint32, mem interfaces.Memory[uint32], inst uint32) {
	rs1 := regs.Get(shared.RS1(inst))

	loadedByte := mem.Load(shared.ImmI(inst) + rs1)

	regs.Set(shared.RD(inst), shared.SignExtend(uint32(loadedByte), 8))
}

// LH (load half-word) loads a 16-bit value from memory at address offset + rs to rd.
// Sign-extending to 32-bits before storing in rd.
//
// lh rd, offset(rs)
func LH(regs interfaces.Register[uint32, types.X32Regs], pc *uint32, mem interfaces.Memory[uint32], inst uint32) {
	rs1 := regs.Get(shared.RS1(inst))

	offsetBase := shared.ImmI(inst) + rs1

	loadedByte0 := mem.Load(offsetBase + 0)
	loadedByte1 := mem.Load(offsetBase + 1)

	loadedHalfW := uint32(loadedByte0) | uint32(loadedByte1)<<8

	regs.Set(shared.RD(inst), shared.SignExtend(loadedHalfW, 16))
}

// LW (load word) loads a 32-bit value from memory at address offset + rs to rd.
//
// lw rd, offset(rs)
func LW(regs interfaces.Register[uint32, types.X32Regs], pc *uint32, mem interfaces.Memory[uint32], inst uint32) {
	rs1 := regs.Get(shared.RS1(inst))

	offsetBase := shared.ImmI(inst) + rs1

	loadedByte0 := mem.Load(offsetBase + 0)
	loadedByte1 := mem.Load(offsetBase + 1)
	loadedByte2 := mem.Load(offsetBase + 2)
	loadedByte3 := mem.Load(offsetBase + 3)

	loadedWord := 0 |
		uint32(loadedByte0)<<(0*8) |
		uint32(loadedByte1)<<(1*8) |
		uint32(loadedByte2)<<(2*8) |
		uint32(loadedByte3)<<(3*8)

	regs.Set(shared.RD(inst), loadedWord)
}

// LBU (load byte unsigned) loads an 8-bit value from memory at address offset + rs to rd.
// Zero-extending to 32-bits before storing in rd.
//
// lbu rd, offset(rs)
func LBU(regs interfaces.Register[uint32, types.X32Regs], pc *uint32, mem interfaces.Memory[uint32], inst uint32) {
	rs1 := regs.Get(shared.RS1(inst))

	loadedByte := mem.Load(shared.ImmI(inst) + rs1)

	regs.Set(shared.RD(inst), uint32(loadedByte))
}

// LHU (load half-word unsigned) loads a 16-bit value from memory at address offset + rs to rd.
// Zero-extending to 32-bits before storing in rd.
//
// lhu rd, offset(rs)
func LHU(regs interfaces.Register[uint32, types.X32Regs], pc *uint32, mem interfaces.Memory[uint32], inst uint32) {
	rs1 := regs.Get(shared.RS1(inst))

	offsetBase := shared.ImmI(inst) + rs1

	loadedByte0 := mem.Load(offsetBase + 0)
	loadedByte1 := mem.Load(offsetBase + 1)

	loadedHalfW := uint32(loadedByte0) | uint32(loadedByte1)<<8

	regs.Set(shared.RD(inst), loadedHalfW)
}

// SB (store byte) stores the lower 8-bits of rs2 to memory at address offset + rs1.
//
// sb rs2, offset(rs1)
func SB(regs interfaces.Register[uint32, types.X32Regs], pc *uint32, mem interfaces.Memory[uint32], inst uint32) {
	rs1 := regs.Get(shared.RS1(inst))
	rs2 := regs.Get(shared.RS2(inst))

	mem.Store(shared.ImmS(inst)+rs1, byte(rs2))
}

// SH (store half-word) stores the lower 16-bits of rs2 to memory at address offset + rs1.
//
// sb rs2, offset(rs1)
func SH(regs interfaces.Register[uint32, types.X32Regs], pc *uint32, mem interfaces.Memory[uint32], inst uint32) {
	rs2 := regs.Get(shared.RS2(inst))

	rs2Value := rs2

	byte0 := byte(rs2Value >> 0)
	byte1 := byte(rs2Value >> 8)

	offsetBase := shared.ImmS(inst) + uint32(shared.RS1(inst))

	mem.Store(offsetBase+0, byte0)
	mem.Store(offsetBase+1, byte1)
}

// SW (store word) stores the 32-bits of rs2 to memory at address offset + rs1.
//
// sb rs2, offset(rs1)
func SW(regs interfaces.Register[uint32, types.X32Regs], pc *uint32, mem interfaces.Memory[uint32], inst uint32) {
	rs2 := regs.Get(shared.RS2(inst))

	rs2Value := rs2

	byte0 := byte(rs2Value >> (0 * 8))
	byte1 := byte(rs2Value >> (1 * 8))
	byte2 := byte(rs2Value >> (2 * 8))
	byte3 := byte(rs2Value >> (3 * 8))

	offsetBase := shared.ImmS(inst) + uint32(shared.RS1(inst))

	mem.Store(offsetBase+0, byte0)
	mem.Store(offsetBase+1, byte1)
	mem.Store(offsetBase+2, byte2)
	mem.Store(offsetBase+3, byte3)
}

// ADDI (add immediate) adds the sign-extended 12-bit immediate to register rs1. Arithmetic overflow is ignored and
// the result is simply the low XLEN bits of the result.
//
// addi rd, rs1, imm
func ADDI(regs interfaces.Register[uint32, types.X32Regs], pc *uint32, mem interfaces.Memory[uint32], inst uint32) {
	rs1 := regs.Get(shared.RS1(inst))

	regs.Set(shared.RD(inst), rs1+shared.ImmI(inst))
}

// SLTI (set less than immediate) places the value 1 in register rd if register rs1 is less than the sign-
// extended immediate when both are treated as signed numbers, else 0 is written to rd.
//
// slti rd, rs1, imm
func SLTI(regs interfaces.Register[uint32, types.X32Regs], pc *uint32, mem interfaces.Memory[uint32], inst uint32) {

}

// SLTIU (set less than immediate unsigned) places the value 1 in register rd if register rs1 is less than the sign-
// extended immediate when both are treated as unsigned numbers, else 0 is written to rd.
// (i.e., the immediate is first sign-extended to XLEN bits then treated as an unsigned number).
//
// sltiu rd, rs1, imm
func SLTIU(regs interfaces.Register[uint32, types.X32Regs], pc *uint32, mem interfaces.Memory[uint32], inst uint32) {

}

// XORI (xor immediate) perform a bitwise XOR on register rs1 and then sign-extended 12-bit immediate and place the result in rd.
//
// xori rd, rs1, imm
func XORI(regs interfaces.Register[uint32, types.X32Regs], pc *uint32, mem interfaces.Memory[uint32], inst uint32) {

}

// ORI (or immediate) perform a bitwise OR on register rs1 and then sign-extended 12-bit immediate and place the result in rd.
//
// ori rd, rs1, imm
func ORI(regs interfaces.Register[uint32, types.X32Regs], pc *uint32, mem interfaces.Memory[uint32], inst uint32) {

}

// ANDI (and immediate) perform a bitwise AND on register rs1 and then sign-extended 12-bit immediate and place the result in rd.
//
// andi rd, rs1, imm
func ANDI(regs interfaces.Register[uint32, types.X32Regs], pc *uint32, mem interfaces.Memory[uint32], inst uint32) {

}

// SLLI (shift left logic immediate) performs a bitwise shift left on register rs1 by the immediate shift ammount and place
// the result in rd. SLLI is a logical left shift (zeros are shifted into the lower bits).
//
// slli rd, rs1, shamt
func SLLI(regs interfaces.Register[uint32, types.X32Regs], pc *uint32, mem interfaces.Memory[uint32], inst uint32) {

}

// SRLI (shift right logic immediate) performs a bitwise shift right on register rs1 by the immediate shift ammount and place
// the result in rd. SRLI is a logical right shift (zeros are shifted into the upper bits).
//
// srli rd, rs1, shamt
func SRLI(regs interfaces.Register[uint32, types.X32Regs], pc *uint32, mem interfaces.Memory[uint32], inst uint32) {

}

// SRAI (shift right arithmetic immediate) performs a bitwise shift right on register rs1 by the immediate shift ammount and place
// the result in rd. SRAI is an arithmetic right shift (the original sign bit is copied into the vacated upper bits).
//
// srai rd, rs1, shamt
func SRAI(regs interfaces.Register[uint32, types.X32Regs], pc *uint32, mem interfaces.Memory[uint32], inst uint32) {

}

// ADD performs the addition of rs1 and rs2. Overflows are ignored and the low XLEN bits of results are written to the
// destination rd.
//
// add rd, rs1, rs2
func ADD(regs interfaces.Register[uint32, types.X32Regs], pc *uint32, mem interfaces.Memory[uint32], inst uint32) {
	rs1 := regs.Get(shared.RS1(inst))
	rs2 := regs.Get(shared.RS2(inst))

	regs.Set(shared.RD(inst), rs1+rs2)
}

// SUB performs the subtraction of rs2 from rs1. Overflows are ignored and the low XLEN bits of results are written to the
// destination rd.
//
// sub rd, rs1, rs2
func SUB(regs interfaces.Register[uint32, types.X32Regs], pc *uint32, mem interfaces.Memory[uint32], inst uint32) {
	rs1 := regs.Get(shared.RS1(inst))
	rs2 := regs.Get(shared.RS2(inst))

	regs.Set(shared.RD(inst), rs1-rs2)
}

// SLL (shift left logical) performs a logical left shift on the value in register rs1 by the shift amount held in the
// lower 5 bits of register rs2 and place the result in rd.
//
// sll rd, rs1, rs2
func SLL(regs interfaces.Register[uint32, types.X32Regs], pc *uint32, mem interfaces.Memory[uint32], inst uint32) {

}

// SLT (set less than) performs signed comparison writing 1 to rd if rs1 < rs2, 0 otherwise.
//
// slt rd, rs1, rs2
func SLT(regs interfaces.Register[uint32, types.X32Regs], pc *uint32, mem interfaces.Memory[uint32], inst uint32) {

}

// SLTU (set less than un signed) performs unsigned comparison writing 1 to rd if rs1 < rs2, 0 otherwise.
//
// sltu rd, rs1, rs2
func SLTU(regs interfaces.Register[uint32, types.X32Regs], pc *uint32, mem interfaces.Memory[uint32], inst uint32) {

}

// XOR performs bitwise logical XOR operation (rd = rs1 XOR rs2).
//
// xor rd, rs1, rs2
func XOR(regs interfaces.Register[uint32, types.X32Regs], pc *uint32, mem interfaces.Memory[uint32], inst uint32) {

}

// SRL (shift right logical) performs a logical right shift on the value in register rs1 by the shift amount held in the lower 5 bits of
// register rs2 and place the result in rd.
//
// srl rd, rs1, rs2
func SRL(regs interfaces.Register[uint32, types.X32Regs], pc *uint32, mem interfaces.Memory[uint32], inst uint32) {

}

// SRA (shift right arithmetic) performs an arithmetic right shift on the value in register rs1 by the shift amount held in the lower 5 bits
// of register rs2 and place the result in rd.
//
// sra rd, rs1, rs2
func SRA(regs interfaces.Register[uint32, types.X32Regs], pc *uint32, mem interfaces.Memory[uint32], inst uint32) {

}

// OR performs bitwise logical OR operation (rd = rs1 OR rs2).
//
// or rd, rs1, rs2
func OR(regs interfaces.Register[uint32, types.X32Regs], pc *uint32, mem interfaces.Memory[uint32], inst uint32) {

}

// AND performs bitwise logical AND operation (rd = rs1 AND rs2).
//
// and rd, rs1, rs2
func AND(regs interfaces.Register[uint32, types.X32Regs], pc *uint32, mem interfaces.Memory[uint32], inst uint32) {

}

func FENCE(regs interfaces.Register[uint32, types.X32Regs], pc *uint32, mem interfaces.Memory[uint32], inst uint32) {

}

// ECALL (environment call) is used to make a service request to the execution environment. The EEI
// will define how parameters for the service request are passed, but usually these will be in defined
// locations in the integer register file.
func ECALL(regs interfaces.Register[uint32, types.X32Regs], pc *uint32, mem interfaces.Memory[uint32], inst uint32) {
	// NOP
}

// EBREAK (environment break) is used to return control to a debugging environment.
func EBREAK(regs interfaces.Register[uint32, types.X32Regs], pc *uint32, mem interfaces.Memory[uint32], inst uint32) {
	// NOP
}
