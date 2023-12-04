package instructionset

import (
	"github.com/d1360-64rc14/risc-v-emulator/pkg/instructionset/rv32ibase"
	"github.com/d1360-64rc14/risc-v-emulator/pkg/interfaces"
	"github.com/d1360-64rc14/risc-v-emulator/pkg/shared"
)

func RV32IBase(regs *[32]uint32, pc *uint32, mem interfaces.Memory[uint32], inst uint32) shared.Instruction[uint32] {
	switch {
	case inst&shared.U_TYPE == rv32ibase.SIGNATURE_LUI:
		return rv32ibase.LUI
	case inst&shared.U_TYPE == rv32ibase.SIGNATURE_AUIPC:
		return rv32ibase.AUIPC
	case inst&shared.J_TYPE == rv32ibase.SIGNATURE_JAL:
		return rv32ibase.JAL
	case inst&shared.I_TYPE == rv32ibase.SIGNATURE_JALR:
		return rv32ibase.JALR
	case inst&shared.B_TYPE == rv32ibase.SIGNATURE_BEQ:
		return rv32ibase.BEQ
	case inst&shared.B_TYPE == rv32ibase.SIGNATURE_BNE:
		return rv32ibase.BNE
	case inst&shared.B_TYPE == rv32ibase.SIGNATURE_BLT:
		return rv32ibase.BLT
	case inst&shared.B_TYPE == rv32ibase.SIGNATURE_BGE:
		return rv32ibase.BGE
	case inst&shared.B_TYPE == rv32ibase.SIGNATURE_BLTU:
		return rv32ibase.BLTU
	case inst&shared.B_TYPE == rv32ibase.SIGNATURE_BGEU:
		return rv32ibase.BGEU
	case inst&shared.I_TYPE == rv32ibase.SIGNATURE_LB:
		return rv32ibase.LB
	case inst&shared.I_TYPE == rv32ibase.SIGNATURE_LH:
		return rv32ibase.LH
	case inst&shared.I_TYPE == rv32ibase.SIGNATURE_LW:
		return rv32ibase.LW
	case inst&shared.I_TYPE == rv32ibase.SIGNATURE_LBU:
		return rv32ibase.LBU
	case inst&shared.I_TYPE == rv32ibase.SIGNATURE_LHU:
		return rv32ibase.LHU
	case inst&shared.S_TYPE == rv32ibase.SIGNATURE_SB:
		return rv32ibase.SB
	case inst&shared.S_TYPE == rv32ibase.SIGNATURE_SH:
		return rv32ibase.SH
	case inst&shared.S_TYPE == rv32ibase.SIGNATURE_SW:
		return rv32ibase.SW
	case inst&shared.I_TYPE == rv32ibase.SIGNATURE_ADDI:
		return rv32ibase.ADDI
	case inst&shared.I_TYPE == rv32ibase.SIGNATURE_SLTI:
		return rv32ibase.SLTI
	case inst&shared.I_TYPE == rv32ibase.SIGNATURE_SLTIU:
		return rv32ibase.SLTIU
	case inst&shared.I_TYPE == rv32ibase.SIGNATURE_XORI:
		return rv32ibase.XORI
	case inst&shared.I_TYPE == rv32ibase.SIGNATURE_ORI:
		return rv32ibase.ORI
	case inst&shared.I_TYPE == rv32ibase.SIGNATURE_ANDI:
		return rv32ibase.ANDI
	case inst&shared.I_TYPE == rv32ibase.SIGNATURE_SLLI:
		return rv32ibase.SLLI
	case inst&shared.I_TYPE == rv32ibase.SIGNATURE_SRLI:
		return rv32ibase.SRLI
	case inst&shared.I_TYPE == rv32ibase.SIGNATURE_SRAI:
		return rv32ibase.SRAI
	case inst&shared.R_TYPE == rv32ibase.SIGNATURE_ADD:
		return rv32ibase.ADD
	case inst&shared.R_TYPE == rv32ibase.SIGNATURE_SUB:
		return rv32ibase.SUB
	case inst&shared.R_TYPE == rv32ibase.SIGNATURE_SLL:
		return rv32ibase.SLL
	case inst&shared.R_TYPE == rv32ibase.SIGNATURE_SLT:
		return rv32ibase.SLT
	case inst&shared.R_TYPE == rv32ibase.SIGNATURE_SLTU:
		return rv32ibase.SLTU
	case inst&shared.R_TYPE == rv32ibase.SIGNATURE_XOR:
		return rv32ibase.XOR
	case inst&shared.R_TYPE == rv32ibase.SIGNATURE_SRL:
		return rv32ibase.SRL
	case inst&shared.R_TYPE == rv32ibase.SIGNATURE_SRA:
		return rv32ibase.SRA
	case inst&shared.R_TYPE == rv32ibase.SIGNATURE_OR:
		return rv32ibase.OR
	case inst&shared.R_TYPE == rv32ibase.SIGNATURE_AND:
		return rv32ibase.AND
	case inst&shared.I_TYPE == rv32ibase.SIGNATURE_FENCE:
		return rv32ibase.FENCE
	case inst&shared.I_TYPE == rv32ibase.SIGNATURE_ECALL:
		return rv32ibase.ECALL
	case inst&shared.I_TYPE == rv32ibase.SIGNATURE_EBREAK:
		return rv32ibase.EBREAK
	}

	return NotFound
}
