package instruction_test

import (
	"testing"

	"github.com/d1360-64rc14/risc-v-emulator/internal/instruction"
)

func TestOpCode(t *testing.T) {
	testCases := []struct {
		desc     string
		inst     uint32
		expected uint32
	}{
		{"only zeros with blank zeros", 0b0000000000000000000000000_0000000, 0b0000000},
		{"only zeros with blank ones", 0b1111111111111111111111111_0000000, 0b0000000},
		{"only ones with blank zeros", 0b0000000000000000000000000_1111111, 0b1111111},
		{"only ones with blank ones", 0b1111111111111111111111111_1111111, 0b1111111},
		{"one alternating with blank zeros", 0b0000000000000000000000000_1010101, 0b1010101},
		{"one alternating with blank ones", 0b1111111111111111111111111_1010101, 0b1010101},
		{"zero alternating with blank zeros", 0b0000000000000000000000000_0101010, 0b0101010},
		{"zero alternating with blank ones", 0b1111111111111111111111111_0101010, 0b0101010},
		{"one incrementing with blank zeros", 0b0000000000000000000000000_1011000, 0b1011000},
		{"one incrementing with blank ones", 0b1111111111111111111111111_1011000, 0b1011000},
		{"zero incrementing with blank zeros", 0b0000000000000000000000000_0001101, 0b0001101},
		{"zero incremenring with blank ones", 0b1111111111111111111111111_0001101, 0b0001101},
	}

	for _, c := range testCases {
		t.Run(c.desc, func(t *testing.T) {
			got := instruction.OpCode(c.inst)

			if got != c.expected {
				t.Errorf("E %b;\nG %b", c.expected, got)
			}
		})
	}
}
func TestRS1(t *testing.T) {
	testCases := []struct {
		desc     string
		inst     uint32
		expected uint32
	}{
		{"only zeros with blank zeros", 0b000000000000_00000_000000000000000, 0b00000},
		{"only zeros with blank ones", 0b111111111111_00000_111111111111111, 0b00000},
		{"only ones with blank zeros", 0b000000000000_11111_000000000000000, 0b11111},
		{"only ones with blank ones", 0b111111111111_11111_111111111111111, 0b11111},
		{"one alternating with blank zeros", 0b000000000000_10101_000000000000000, 0b10101},
		{"one alternating with blank ones", 0b111111111111_10101_111111111111111, 0b10101},
		{"zero alternating with blank zeros", 0b000000000000_01010_000000000000000, 0b01010},
		{"zero alternating with blank ones", 0b111111111111_01010_111111111111111, 0b01010},
		{"one incrementing with blank zeros", 0b000000000000_10110_000000000000000, 0b10110},
		{"one incrementing with blank ones", 0b111111111111_10110_111111111111111, 0b10110},
		{"zero incrementing with blank zeros", 0b000000000000_01101_000000000000000, 0b01101},
		{"zero incremenring with blank ones", 0b111111111111_01101_111111111111111, 0b01101},
	}

	for _, c := range testCases {
		t.Run(c.desc, func(t *testing.T) {
			got := instruction.RS1(c.inst)

			if got != c.expected {
				t.Errorf("E %b;\nG %b", c.expected, got)
			}
		})
	}
}
func TestRS2(t *testing.T) {
	testCases := []struct {
		desc     string
		inst     uint32
		expected uint32
	}{
		{"only zeros with blank zeros", 0b0000000_00000_00000000000000000000, 0b00000},
		{"only zeros with blank ones", 0b1111111_00000_11111111111111111111, 0b00000},
		{"only ones with blank zeros", 0b0000000_11111_00000000000000000000, 0b11111},
		{"only ones with blank ones", 0b1111111_11111_11111111111111111111, 0b11111},
		{"one alternating with blank zeros", 0b0000000_10101_00000000000000000000, 0b10101},
		{"one alternating with blank ones", 0b1111111_10101_11111111111111111111, 0b10101},
		{"zero alternating with blank zeros", 0b0000000_01010_00000000000000000000, 0b01010},
		{"zero alternating with blank ones", 0b1111111_01010_11111111111111111111, 0b01010},
		{"one incrementing with blank zeros", 0b0000000_10110_00000000000000000000, 0b10110},
		{"one incrementing with blank ones", 0b1111111_10110_11111111111111111111, 0b10110},
		{"zero incrementing with blank zeros", 0b0000000_01101_00000000000000000000, 0b01101},
		{"zero incremenring with blank ones", 0b1111111_01101_11111111111111111111, 0b01101},
	}

	for _, c := range testCases {
		t.Run(c.desc, func(t *testing.T) {
			got := instruction.RS2(c.inst)

			if got != c.expected {
				t.Errorf("E %b;\nG %b", c.expected, got)
			}
		})
	}
}
func TestRD(t *testing.T) {
	testCases := []struct {
		desc     string
		inst     uint32
		expected uint32
	}{
		{"only zeros with blank zeros", 0b00000000000000000000_00000_0000000, 0b00000},
		{"only zeros with blank ones", 0b11111111111111111111_00000_1111111, 0b00000},
		{"only ones with blank zeros", 0b00000000000000000000_11111_0000000, 0b11111},
		{"only ones with blank ones", 0b11111111111111111111_11111_1111111, 0b11111},
		{"one alternating with blank zeros", 0b00000000000000000000_10101_0000000, 0b10101},
		{"one alternating with blank ones", 0b11111111111111111111_10101_1111111, 0b10101},
		{"zero alternating with blank zeros", 0b00000000000000000000_01010_0000000, 0b01010},
		{"zero alternating with blank ones", 0b11111111111111111111_01010_1111111, 0b01010},
		{"one incrementing with blank zeros", 0b00000000000000000000_10110_0000000, 0b10110},
		{"one incrementing with blank ones", 0b11111111111111111111_10110_1111111, 0b10110},
		{"zero incrementing with blank zeros", 0b00000000000000000000_01101_0000000, 0b01101},
		{"zero incremenring with blank ones", 0b11111111111111111111_01101_1111111, 0b01101},
	}

	for _, c := range testCases {
		t.Run(c.desc, func(t *testing.T) {
			got := instruction.RD(c.inst)

			if got != c.expected {
				t.Errorf("E %b;\nG %b", c.expected, got)
			}
		})
	}
}
func TestFunct3(t *testing.T) {
	testCases := []struct {
		desc     string
		inst     uint32
		expected uint32
	}{
		{"only zeros with blank zeros", 0b00000000000000000_000_000000000000, 0b000},
		{"only zeros with blank ones", 0b11111111111111111_000_111111111111, 0b000},
		{"only ones with blank zeros", 0b00000000000000000_111_000000000000, 0b111},
		{"only ones with blank ones", 0b11111111111111111_111_111111111111, 0b111},
		{"one alternating with blank zeros", 0b00000000000000000_101_000000000000, 0b101},
		{"one alternating with blank ones", 0b11111111111111111_101_111111111111, 0b101},
		{"zero alternating with blank zeros", 0b00000000000000000_010_000000000000, 0b010},
		{"zero alternating with blank ones", 0b11111111111111111_010_111111111111, 0b010},
		{"one incrementing with blank zeros", 0b00000000000000000_100_000000000000, 0b100},
		{"one incrementing with blank ones", 0b11111111111111111_100_111111111111, 0b100},
		{"zero incrementing with blank zeros", 0b00000000000000000_001_000000000000, 0b001},
		{"zero incremenring with blank ones", 0b11111111111111111_001_111111111111, 0b001},
	}

	for _, c := range testCases {
		t.Run(c.desc, func(t *testing.T) {
			got := instruction.Funct3(c.inst)

			if got != c.expected {
				t.Errorf("E %b;\nG %b", c.expected, got)
			}
		})
	}
}
func TestFunct7(t *testing.T) {
	testCases := []struct {
		desc     string
		inst     uint32
		expected uint32
	}{
		{"only zeros with blank zeros", 0b0000000_0000000000000000000000000, 0b0000000},
		{"only zeros with blank ones", 0b0000000_1111111111111111111111111, 0b0000000},
		{"only ones with blank zeros", 0b1111111_0000000000000000000000000, 0b1111111},
		{"only ones with blank ones", 0b1111111_1111111111111111111111111, 0b1111111},
		{"one alternating with blank zeros", 0b1010101_0000000000000000000000000, 0b1010101},
		{"one alternating with blank ones", 0b1010101_1111111111111111111111111, 0b1010101},
		{"zero alternating with blank zeros", 0b0101010_0000000000000000000000000, 0b0101010},
		{"zero alternating with blank ones", 0b0101010_1111111111111111111111111, 0b0101010},
		{"one incrementing with blank zeros", 0b1011000_0000000000000000000000000, 0b1011000},
		{"one incrementing with blank ones", 0b1011000_1111111111111111111111111, 0b1011000},
		{"zero incrementing with blank zeros", 0b0001101_0000000000000000000000000, 0b0001101},
		{"zero incremenring with blank ones", 0b0001101_1111111111111111111111111, 0b0001101},
	}

	for _, c := range testCases {
		t.Run(c.desc, func(t *testing.T) {
			got := instruction.Funct7(c.inst)

			if got != c.expected {
				t.Errorf("E %b;\nG %b", c.expected, got)
			}
		})
	}
}

func TestSignExtend(t *testing.T) {
	testCases := []struct {
		desc     string
		number   uint32
		bitSize  uint32
		expected uint32
	}{
		{"bitsize 1 with sign 1", 0b1, 1, 0b1111111111111111111111111111111_1},
		{"bitsize 1 with sign 0", 0b0, 1, 0b0000000000000000000000000000000_0},
		{"bitsize 6 with sign 1", 0b100101, 6, 0b11111111111111111111111111_100101},
		{"bitsize 6 with sign 0", 0b000101, 6, 0b00000000000000000000000000_000101},
		{"bitsize 8 with sign 1", 0b10000101, 8, 0b111111111111111111111111_10000101},
		{"bitsize 8 with sign 0", 0b00000101, 8, 0b000000000000000000000000_00000101},
	}

	for _, c := range testCases {
		t.Run(c.desc, func(t *testing.T) {
			got := instruction.SignExtend(c.number, c.bitSize)

			if got != c.expected {
				t.Errorf("E %b;\nG %b", c.expected, got)
			}
		})
	}
}

func TestImmI(t *testing.T) {
	testCases := []struct {
		desc     string
		inst     uint32
		expected uint32
	}{
		{"only zeros with blank zeros", 0b000000000000_00000_000_00000_0000000, 0b00000000000000000000_000000000000},
		{"only zeros with blank ones", 0b000000000000_11111_111_11111_1111111, 0b00000000000000000000_000000000000},
		{"only ones with blank zeros", 0b111111111111_00000_000_00000_0000000, 0b11111111111111111111_111111111111},
		{"only one with blank ones", 0b111111111111_11111_111_11111_1111111, 0b11111111111111111111_111111111111},
		{"one alternating with blank zeros", 0b101010101010_00000_000_00000_0000000, 0b11111111111111111111_101010101010},
		{"one alternating with blank ones", 0b101010101010_11111_111_11111_1111111, 0b11111111111111111111_101010101010},
		{"zero alernating with blank zeros", 0b010101010101_00000_000_00000_0000000, 0b00000000000000000000_010101010101},
		{"zero alernating with blank ones", 0b010101010101_11111_111_11111_1111111, 0b00000000000000000000_010101010101},
		{"one incrementing with blank zeros", 0b101100111000_00000_000_00000_0000000, 0b11111111111111111111_101100111000},
		{"one incrementing with blank ones", 0b101100111000_11111_111_11111_1111111, 0b11111111111111111111_101100111000},
		{"zero incrementing with blank zeros", 0b000111001101_00000_000_00000_0000000, 0b00000000000000000000_000111001101},
		{"zero incrementing with blank ones", 0b000111001101_11111_111_11111_1111111, 0b00000000000000000000_000111001101},
	}

	for _, c := range testCases {
		t.Run(c.desc, func(t *testing.T) {
			got := instruction.ImmI(c.inst)

			if got != c.expected {
				t.Errorf("E %b;\nG %b", c.expected, got)
			}
		})
	}
}

func TestImmS(t *testing.T) {
	testCases := []struct {
		desc     string
		inst     uint32
		expected uint32
	}{
		{"only zeros with blank zeros", 0b0000000_00000_00000_000_00000_0000000, 0b00000000000000000000_000000000000},
		{"only zeros with blank ones", 0b0000000_11111_11111_111_00000_1111111, 0b00000000000000000000_000000000000},
		{"only ones with blank zeros", 0b1111111_00000_00000_000_11111_0000000, 0b11111111111111111111_111111111111},
		{"only ones with blank ones", 0b1111111_11111_11111_111_11111_1111111, 0b11111111111111111111_111111111111},
		{"one alternating with blank zeros", 0b1010101_00000_00000_000_01010_0000000, 0b11111111111111111111_101010101010},
		{"one alternating with blank ones", 0b1010101_11111_11111_111_01010_1111111, 0b11111111111111111111_101010101010},
		{"zero alernating with blank zeros", 0b0101010_00000_00000_000_10101_0000000, 0b00000000000000000000_010101010101},
		{"zero alernating with blank ones", 0b0101010_11111_11111_111_10101_1111111, 0b00000000000000000000_010101010101},
		{"one incrementing with blank zeros", 0b1011001_00000_00000_000_11000_0000000, 0b11111111111111111111_101100111000},
		{"one incrementing with blank ones", 0b1011001_11111_11111_111_11000_1111111, 0b11111111111111111111_101100111000},
		{"zero incrementing with blank zeros", 0b0001110_00000_00000_000_01101_0000000, 0b00000000000000000000_000111001101},
		{"zero incrementing with blank ones", 0b0001110_11111_11111_111_01101_1111111, 0b00000000000000000000_000111001101},
	}

	for _, c := range testCases {
		t.Run(c.desc, func(t *testing.T) {
			got := instruction.ImmS(c.inst)

			if got != c.expected {
				t.Errorf("E %b;\nG %b", c.expected, got)
			}
		})
	}
}

func TestImmB(t *testing.T) {
	testCases := []struct {
		desc     string
		inst     uint32
		expected uint32
	}{
		{"only zeros with blank zeros", 0b0000000_00000_00000_000_00000_0000000, 0b0000000000000000000_0000000000000},
		{"only zeros with blank ones", 0b0000000_11111_11111_111_00000_1111111, 0b0000000000000000000_0000000000000},
		{"only ones with blank zeros", 0b1111111_00000_00000_000_11111_0000000, 0b1111111111111111111_1111111111110},
		{"only ones with blank ones", 0b1111111_11111_11111_111_11111_1111111, 0b1111111111111111111_1111111111110},
		{"one alternating with blank zeros", 0b1101010_00000_00000_000_10100_0000000, 0b1111111111111111111_1010101010100},
		{"one alternating with blank ones", 0b1101010_11111_11111_111_10100_1111111, 0b1111111111111111111_1010101010100},
		{"zero alernating with blank zeros", 0b0010101_00000_00000_000_01011_0000000, 0b0000000000000000000_0101010101010},
		{"zero alernating with blank ones", 0b0010101_11111_11111_111_01011_1111111, 0b0000000000000000000_0101010101010},
		{"one incrementing with blank zeros", 0b1110011_00000_00000_000_10000_0000000, 0b1111111111111111111_1011001110000},
		{"one incrementing with blank ones", 0b1110011_11111_11111_111_10000_1111111, 0b1111111111111111111_1011001110000},
		{"zero incrementing with blank zeros", 0b0001110_00000_00000_000_01100_0000000, 0b0000000000000000000_0000111001100},
		{"zero incrementing with blank ones", 0b0001110_11111_11111_111_01100_1111111, 0b0000000000000000000_0000111001100},
	}

	for _, c := range testCases {
		t.Run(c.desc, func(t *testing.T) {
			got := instruction.ImmB(c.inst)

			if got != c.expected {
				t.Errorf("E %b;\nG %b", c.expected, got)
			}
		})
	}
}

func TestImmU(t *testing.T) {
	testCases := []struct {
		desc     string
		inst     uint32
		expected uint32
	}{
		{"only zeros with blank zeros", 0b00000000000000000000_00000_0000000, 0b00000000000000000000_000000000000},
		{"only zeros with blank ones", 0b00000000000000000000_11111_1111111, 0b00000000000000000000_000000000000},
		{"only ones with blank zeros", 0b11111111111111111111_00000_0000000, 0b11111111111111111111_000000000000},
		{"only ones with blank ones", 0b11111111111111111111_11111_1111111, 0b11111111111111111111_000000000000},
		{"one alternating with blank zeros", 0b10101010101010101010_00000_0000000, 0b10101010101010101010_000000000000},
		{"one alternating with blank ones", 0b10101010101010101010_11111_1111111, 0b10101010101010101010_000000000000},
		{"zero alernating with blank zeros", 0b01010101010101010101_00000_0000000, 0b01010101010101010101_000000000000},
		{"zero alernating with blank ones", 0b01010101010101010101_11111_1111111, 0b01010101010101010101_000000000000},
		{"one incrementing with blank zeros", 0b10110011100011110000_00000_0000000, 0b10110011100011110000_000000000000},
		{"one incrementing with blank ones", 0b10110011100011110000_11111_1111111, 0b10110011100011110000_000000000000},
		{"zero incrementing with blank zeros", 0b00001111000111001101_00000_0000000, 0b00001111000111001101_000000000000},
		{"zero incrementing with blank ones", 0b00001111000111001101_11111_1111111, 0b00001111000111001101_000000000000},
	}

	for _, c := range testCases {
		t.Run(c.desc, func(t *testing.T) {
			got := instruction.ImmU(c.inst)

			if got != c.expected {
				t.Errorf("E %b;\nG %b", c.expected, got)
			}
		})
	}
}

func TestImmJ(t *testing.T) {
	testCases := []struct {
		desc     string
		inst     uint32
		expected uint32
	}{
		{"only zeros with blank zeros", 0b00000000000000000000_00000_0000000, 0b00000000000_000000000000000000000},
		{"only zeros with blank ones", 0b00000000000000000000_11111_1111111, 0b00000000000_000000000000000000000},
		{"only ones with blank zeros", 0b11111111111111111111_00000_0000000, 0b11111111111_111111111111111111110},
		{"only ones with blank ones", 0b11111111111111111111_11111_1111111, 0b11111111111_111111111111111111110},
		{"one alternating with blank zeros", 0b11010101010001010101_00000_0000000, 0b11111111111_101010101010101010100},
		{"one alternating with blank ones", 0b11010101010001010101_11111_1111111, 0b11111111111_101010101010101010100},
		{"zero alernating with blank zeros", 0b00101010101110101010_00000_0000000, 0b00000000000_010101010101010101010},
		{"zero alernating with blank ones", 0b00101010101110101010_11111_1111111, 0b00000000000_010101010101010101010},
		{"one incrementing with blank zeros", 0b10011110000001100111_00000_0000000, 0b11111111111_101100111000111100000},
		{"one incrementing with blank ones", 0b10011110000001100111_11111_1111111, 0b11111111111_101100111000111100000},
		{"zero incrementing with blank zeros", 0b00011100110000001111_00000_0000000, 0b00000000000_000001111000111001100},
		{"zero incrementing with blank ones", 0b00011100110000001111_11111_1111111, 0b00000000000_000001111000111001100},
	}

	for _, c := range testCases {
		t.Run(c.desc, func(t *testing.T) {
			got := instruction.ImmJ(c.inst)

			if got != c.expected {
				t.Errorf("E %b;\nG %b", c.expected, got)
			}
		})
	}
}
