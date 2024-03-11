package rv32i

import (
	"github.com/d1360-64rc14/risc-v-emulator/pkg/instructionset"
	"github.com/d1360-64rc14/risc-v-emulator/pkg/instructionset/rv32i"
	"github.com/d1360-64rc14/risc-v-emulator/pkg/interfaces"
	"github.com/d1360-64rc14/risc-v-emulator/pkg/register/x32register"
	"github.com/d1360-64rc14/risc-v-emulator/pkg/shared"
	"github.com/d1360-64rc14/risc-v-emulator/pkg/types"
)

type RV32I struct {
	regs *x32register.X32Register[uint32]
	pc   uint32
	mem  interfaces.Memory[uint32]
}

func New(
	regs *x32register.X32Register[uint32],
	mem interfaces.Memory[uint32],
) *RV32I {
	return &RV32I{
		regs: regs,
		pc:   0,
		mem:  mem,
	}
}

// Processor loop (fetch, decode, execute, [store])
func (r *RV32I) Start() {
	for {
		inst := r.fetch()
		instFn := r.decode(inst)
		r.execute(instFn, inst)

		r.pc += 4

		if inst&shared.E_TYPE == rv32i.SIGNATURE_EBREAK {
			return
		}
	}
}

func (r *RV32I) fetch() uint32 {
	return 0 |
		uint32(r.mem.Load(r.pc+0))<<(8*0) |
		uint32(r.mem.Load(r.pc+1))<<(8*1) |
		uint32(r.mem.Load(r.pc+2))<<(8*2) |
		uint32(r.mem.Load(r.pc+3))<<(8*3)
}

func (r *RV32I) decode(inst uint32) shared.Instruction[uint32, types.X32Regs] {
	return instructionset.RV32IBase(inst)
}

func (r *RV32I) execute(instFn shared.Instruction[uint32, types.X32Regs], inst uint32) {
	instFn(r.regs, &r.pc, r.mem, inst)
}
