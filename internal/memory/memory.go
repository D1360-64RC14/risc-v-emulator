package memory

type Memory struct {
	data []byte
}

func NewMemory(size int) *Memory {
	return &Memory{
		data: make([]byte, size),
	}
}

func (m *Memory) Load(addr uint32) byte {
	return m.data[addr%uint32(len(m.data))]
}

func (m *Memory) Store(addr uint32, data byte) {
	m.data[addr%uint32(len(m.data))] = data
}
