package puzzle

type DotStatus int

const (
	Unknown DotStatus = iota
	Filled
	Empty
)

type Dot struct {
	status DotStatus
	n      int
}

func NewDot() *Dot {
	return &Dot{
		status: Unknown,
		n:      0,
	}
}
