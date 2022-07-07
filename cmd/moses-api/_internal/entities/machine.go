package entities

type MachineType uint64

const (
	_ MachineType = iota
	IMG
	JinFeng
)

func (m MachineType) String() string {
	switch m {
	case IMG:
		return "IMG MEC-5000"
	case JinFeng:
		return "JinFeng PRG-600"
	default:
		return "Unknown"
	}
}
