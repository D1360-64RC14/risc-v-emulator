package shared

import (
	"math"

	"github.com/d1360-64rc14/risc-v-emulator/refact/interfaces"
	"golang.org/x/exp/constraints"
)

const (
	BIT_SIZE_OPCODE = 7
	BIT_SIZE_RS1    = 5
	BIT_SIZE_RS2    = 5
	BIT_SIZE_RD     = 5
	BIT_SIZE_FUNCT3 = 3
	BIT_SIZE_FUNCT7 = 7
)

type Instruction[A constraints.Unsigned] func(regs *[32]A, pc *A, mem interfaces.Memory[A], inst A)

// OpCode of instruction.
//
// inst[6:0]
func OpCode(inst uint32) uint32 {
	return inst & 0x7F
}

// Register Source 1 of instruction.
//
// inst[19:15]
func RS1(inst uint32) uint32 {
	return inst >> 15 & 0x1F
}

// Register Source 2 of instruction.
//
// inst[24:20]
func RS2(inst uint32) uint32 {
	return inst >> 20 & 0x1F
}

// Register Destination of instruction.
//
// inst[11:7]
func RD(inst uint32) uint32 {
	return inst >> 7 & 0x1F
}

// Funct3 of instruction.
//
// inst[14:12]
func Funct3(inst uint32) uint32 {
	return inst >> 12 & 0x7
}

// Funct7 of instruction.
//
// inst[31:25]
func Funct7(inst uint32) uint32 {
	return inst >> 25
}

// I-immediate of instruction
//
// —-inst[31]-— inst[30:25] inst[24:21] inst[20]
func ImmI(inst uint32) uint32 {
	// --inst[31]-- inst[30:20]
	imm := inst >> 20

	return SignExtend(imm, 12)
}

// S-immediate of instruction
//
// —-inst[31]-— inst[30:25] inst[11:8] inst[7]
func ImmS(inst uint32) uint32 {
	// --inst[31]-- inst[30:25] inst[11:7]
	//
	imm := inst>>7&0x1F | (inst>>25&0x7F)<<5

	return SignExtend(imm, 12)
}

// B-immediate of instruction
//
// —-inst[31]-— inst[7] inst[30:25] inst[11:8] 0
func ImmB(inst uint32) uint32 {
	imm := 0 | (inst>>8&0xF)<<1 | (inst>>25&0x3F)<<5 | (inst>>7&1)<<11 | (inst>>31)<<12

	return SignExtend(imm, 13)
}

// U-immediate of instruction
//
// inst[31] inst[30:20] inst[19:12] --0--
func ImmU(inst uint32) uint32 {
	// inst[31:12] --0--
	return inst & (0xFFFFF << 12)
}

// J-immediate of instruction
//
// —-inst[31]-— inst[19:12] inst[20] inst[30:25] inst[24:21] 0
func ImmJ(inst uint32) uint32 {
	// --inst[31]-- inst[19:12] inst[20] inst[30:21] 0
	imm := 0 | (inst>>21&0x3FF)<<1 | (inst>>20&1)<<11 | (inst>>12&0xFF)<<12 | (inst>>31)<<20

	return SignExtend(imm, 21)
}

// SignExtend replicates the most significant bit to the upper bits of uint32.
//
// Input:
//   - number: 00000000 00000000 00000000 00100101
//   - bitSize: 6
//
// Sign bit is 1 in this example, so it will be extended.
//
// Output: 11111111 11111111 11111111 11100101
func SignExtend(number, bitSize uint32) uint32 {
	signBit := number >> (bitSize - 1) & 1

	return number | math.MaxUint32*signBit<<bitSize
}
