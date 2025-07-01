package facade

import "fmt"

// Subsystem components

type CPU struct{}

func (c *CPU) Freeze() { fmt.Println("CPU: freezing registers") }
func (c *CPU) Jump(position int) {
	fmt.Printf("CPU: jump to address %X\n", position)
}
func (c *CPU) Execute() { fmt.Println("CPU: executing") }

type Memory struct{}

func (m *Memory) Load(position int, data []byte) {
	fmt.Printf("Memory: loading data to address %X\n", position)
}

type HardDrive struct{}

func (h *HardDrive) Read(lba, size int) []byte {
	fmt.Printf("HardDrive: reading %d bytes from LBA %d\n", size, lba)
	return []byte{0xDE, 0xAD, 0xBE, 0xEF}
}

// Facade

type ComputerFacade struct {
	cpu       *CPU
	memory    *Memory
	hardDrive *HardDrive
}

func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{
		cpu:       &CPU{},
		memory:    &Memory{},
		hardDrive: &HardDrive{},
	}
}

func (f *ComputerFacade) Start() {
	const bootAddress = 0x1000
	const bootSector = 40
	const sectorSize = 512

	fmt.Println("Facade: starting computer...")
	f.cpu.Freeze()
	data := f.hardDrive.Read(bootSector, sectorSize)
	f.memory.Load(bootAddress, data)
	f.cpu.Jump(bootAddress)
	f.cpu.Execute()
}

func RunComputerFacadeDemo() {
	computer := NewComputerFacade()
	computer.Start()
}
