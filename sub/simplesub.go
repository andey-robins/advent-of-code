package sub

// Simple sub is a Submarine without the aim based
// movement. Used in Day 2 Part 1
type SimpleSubmarine struct {
	depth    int
	position int
}

func NewSimpleSub() *SimpleSubmarine {
	return &SimpleSubmarine{depth: 0, position: 0}
}

func (s *SimpleSubmarine) Dive(down int) {
	s.depth += down
}

func (s *SimpleSubmarine) Rise(up int) {
	s.depth -= up
}

func (s *SimpleSubmarine) Forward(dist int) {
	s.position += dist
}

func (s *SimpleSubmarine) GetDepth() int {
	return s.depth
}

func (s *SimpleSubmarine) GetPosition() int {
	return s.position
}
