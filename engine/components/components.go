package components

type CompSig uint64

// create unique bit position for each component (up to 64 component)
const (
	PositionSig CompSig = 1 << iota // 1
	VelocitySig                     // 2
)

type Position struct {
	X, Y float64
}

type Velocity struct {
	DX, DY float64
}
