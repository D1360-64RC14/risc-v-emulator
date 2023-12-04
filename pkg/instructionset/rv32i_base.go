package instructionset

import (
	"github.com/d1360-64rc14/risc-v-emulator/pkg/instructionset/rv32i"
	"github.com/d1360-64rc14/risc-v-emulator/pkg/interfaces"
	"github.com/d1360-64rc14/risc-v-emulator/pkg/shared"
)

func RV32IBase(regs *[32]uint32, pc *uint32, mem interfaces.Memory[uint32], inst uint32) shared.Instruction[uint32] {
	switch {
	case inst&shared.U_TYPE == rv32i.SIGNATURE_LUI:
		return rv32i.LUI
	case inst&shared.U_TYPE == rv32i.SIGNATURE_AUIPC:
		return rv32i.AUIPC
	case inst&shared.J_TYPE == rv32i.SIGNATURE_JAL:
		return rv32i.JAL
	case inst&shared.I_TYPE == rv32i.SIGNATURE_JALR:
		return rv32i.JALR
	case inst&shared.B_TYPE == rv32i.SIGNATURE_BEQ:
		return rv32i.BEQ
	case inst&shared.B_TYPE == rv32i.SIGNATURE_BNE:
		return rv32i.BNE
	case inst&shared.B_TYPE == rv32i.SIGNATURE_BLT:
		return rv32i.BLT
	case inst&shared.B_TYPE == rv32i.SIGNATURE_BGE:
		return rv32i.BGE
	case inst&shared.B_TYPE == rv32i.SIGNATURE_BLTU:
		return rv32i.BLTU
	case inst&shared.B_TYPE == rv32i.SIGNATURE_BGEU:
		return rv32i.BGEU
	case inst&shared.I_TYPE == rv32i.SIGNATURE_LB:
		return rv32i.LB
	case inst&shared.I_TYPE == rv32i.SIGNATURE_LH:
		return rv32i.LH
	case inst&shared.I_TYPE == rv32i.SIGNATURE_LW:
		return rv32i.LW
	case inst&shared.I_TYPE == rv32i.SIGNATURE_LBU:
		return rv32i.LBU
	case inst&shared.I_TYPE == rv32i.SIGNATURE_LHU:
		return rv32i.LHU
	case inst&shared.S_TYPE == rv32i.SIGNATURE_SB:
		return rv32i.SB
	case inst&shared.S_TYPE == rv32i.SIGNATURE_SH:
		return rv32i.SH
	case inst&shared.S_TYPE == rv32i.SIGNATURE_SW:
		return rv32i.SW
	case inst&shared.I_TYPE == rv32i.SIGNATURE_ADDI:
		return rv32i.ADDI
	case inst&shared.I_TYPE == rv32i.SIGNATURE_SLTI:
		return rv32i.SLTI
	case inst&shared.I_TYPE == rv32i.SIGNATURE_SLTIU:
		return rv32i.SLTIU
	case inst&shared.I_TYPE == rv32i.SIGNATURE_XORI:
		return rv32i.XORI
	case inst&shared.I_TYPE == rv32i.SIGNATURE_ORI:
		return rv32i.ORI
	case inst&shared.I_TYPE == rv32i.SIGNATURE_ANDI:
		return rv32i.ANDI
	case inst&shared.I_TYPE == rv32i.SIGNATURE_SLLI:
		return rv32i.SLLI
	case inst&shared.I_TYPE == rv32i.SIGNATURE_SRLI:
		return rv32i.SRLI
	case inst&shared.I_TYPE == rv32i.SIGNATURE_SRAI:
		return rv32i.SRAI
	case inst&shared.R_TYPE == rv32i.SIGNATURE_ADD:
		return rv32i.ADD
	case inst&shared.R_TYPE == rv32i.SIGNATURE_SUB:
		return rv32i.SUB
	case inst&shared.R_TYPE == rv32i.SIGNATURE_SLL:
		return rv32i.SLL
	case inst&shared.R_TYPE == rv32i.SIGNATURE_SLT:
		return rv32i.SLT
	case inst&shared.R_TYPE == rv32i.SIGNATURE_SLTU:
		return rv32i.SLTU
	case inst&shared.R_TYPE == rv32i.SIGNATURE_XOR:
		return rv32i.XOR
	case inst&shared.R_TYPE == rv32i.SIGNATURE_SRL:
		return rv32i.SRL
	case inst&shared.R_TYPE == rv32i.SIGNATURE_SRA:
		return rv32i.SRA
	case inst&shared.R_TYPE == rv32i.SIGNATURE_OR:
		return rv32i.OR
	case inst&shared.R_TYPE == rv32i.SIGNATURE_AND:
		return rv32i.AND
	case inst&shared.I_TYPE == rv32i.SIGNATURE_FENCE:
		return rv32i.FENCE
	case inst&shared.I_TYPE == rv32i.SIGNATURE_ECALL:
		return rv32i.ECALL
	case inst&shared.I_TYPE == rv32i.SIGNATURE_EBREAK:
		return rv32i.EBREAK
	}

	return NotFound
}
