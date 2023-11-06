package riscv

import (
	"github.com/d1360-64rc14/risc-v-emulator/internal/instruction"
	"github.com/d1360-64rc14/risc-v-emulator/internal/memory"
)

type RiscV32 struct {
	regs [32]uint32
	pc   uint32
	mem  *memory.Memory
}

func NewRiscV32(mem *memory.Memory) *RiscV32 {
	return &RiscV32{
		regs: [32]uint32{},
		pc:   0,
		mem:  mem,
	}
}

// Processor loop (fetch, decode, execute, [store])
func (r *RiscV32) Run() {
	for {
		inst := r.fetch()
		instFn := r.decode(inst)
		r.execute(instFn, inst)

		r.pc += 4
		r.regs[ZERO] = 0
	}
}

func (r *RiscV32) fetch() uint32 {
	return 0 |
		uint32(r.mem.Load(r.pc+0))<<8*0 |
		uint32(r.mem.Load(r.pc+1))<<8*1 |
		uint32(r.mem.Load(r.pc+2))<<8*2 |
		uint32(r.mem.Load(r.pc+3))<<8*3
}

func (r *RiscV32) decode(inst uint32) instruction.Instruction {
	switch {
	case inst&U_TYPE == instruction.SIGNATURE_LUI:
		return instruction.LUI
	case inst&U_TYPE == instruction.SIGNATURE_AUIPC:
		return instruction.AUIPC
	case inst&J_TYPE == instruction.SIGNATURE_JAL:
		return instruction.JAL
	case inst&I_TYPE == instruction.SIGNATURE_JALR:
		return instruction.JALR
	case inst&B_TYPE == instruction.SIGNATURE_BEQ:
		return instruction.BEQ
	case inst&B_TYPE == instruction.SIGNATURE_BNE:
		return instruction.BNE
	case inst&B_TYPE == instruction.SIGNATURE_BLT:
		return instruction.BLT
	case inst&B_TYPE == instruction.SIGNATURE_BGE:
		return instruction.BGE
	case inst&B_TYPE == instruction.SIGNATURE_BLTU:
		return instruction.BLTU
	case inst&B_TYPE == instruction.SIGNATURE_BGEU:
		return instruction.BGEU
	case inst&I_TYPE == instruction.SIGNATURE_LB:
		return instruction.LB
	case inst&I_TYPE == instruction.SIGNATURE_LH:
		return instruction.LH
	case inst&I_TYPE == instruction.SIGNATURE_LW:
		return instruction.LW
	case inst&I_TYPE == instruction.SIGNATURE_LBU:
		return instruction.LBU
	case inst&I_TYPE == instruction.SIGNATURE_LHU:
		return instruction.LHU
	case inst&S_TYPE == instruction.SIGNATURE_SB:
		return instruction.SB
	case inst&S_TYPE == instruction.SIGNATURE_SH:
		return instruction.SH
	case inst&S_TYPE == instruction.SIGNATURE_SW:
		return instruction.SW
	case inst&I_TYPE == instruction.SIGNATURE_ADDI:
		return instruction.ADDI
	case inst&I_TYPE == instruction.SIGNATURE_SLTI:
		return instruction.SLTI
	case inst&I_TYPE == instruction.SIGNATURE_SLTIU:
		return instruction.SLTIU
	case inst&I_TYPE == instruction.SIGNATURE_XORI:
		return instruction.XORI
	case inst&I_TYPE == instruction.SIGNATURE_ORI:
		return instruction.ORI
	case inst&I_TYPE == instruction.SIGNATURE_ANDI:
		return instruction.ANDI
	case inst&I_TYPE == instruction.SIGNATURE_SLLI:
		return instruction.SLLI
	case inst&I_TYPE == instruction.SIGNATURE_SRLI:
		return instruction.SRLI
	case inst&I_TYPE == instruction.SIGNATURE_SRAI:
		return instruction.SRAI
	case inst&R_TYPE == instruction.SIGNATURE_ADD:
		return instruction.ADD
	case inst&R_TYPE == instruction.SIGNATURE_SUB:
		return instruction.SUB
	case inst&R_TYPE == instruction.SIGNATURE_SLL:
		return instruction.SLL
	case inst&R_TYPE == instruction.SIGNATURE_SLT:
		return instruction.SLT
	case inst&R_TYPE == instruction.SIGNATURE_SLTU:
		return instruction.SLTU
	case inst&R_TYPE == instruction.SIGNATURE_XOR:
		return instruction.XOR
	case inst&R_TYPE == instruction.SIGNATURE_SRL:
		return instruction.SRL
	case inst&R_TYPE == instruction.SIGNATURE_SRA:
		return instruction.SRA
	case inst&R_TYPE == instruction.SIGNATURE_OR:
		return instruction.OR
	case inst&R_TYPE == instruction.SIGNATURE_AND:
		return instruction.AND
	case inst&I_TYPE == instruction.SIGNATURE_FENCE:
		return instruction.FENCE
	case inst&I_TYPE == instruction.SIGNATURE_ECALL:
		return instruction.ECALL
	case inst&I_TYPE == instruction.SIGNATURE_EBREAK:
		return instruction.EBREAK
	}

	return instruction.NotFound
}

// Store step is included in execute.
func (r *RiscV32) execute(instFn instruction.Instruction, inst uint32) {
	instFn(r.regs, &r.pc, r.mem, inst)
}
