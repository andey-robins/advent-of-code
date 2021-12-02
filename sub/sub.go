package sub

// Submarine is the ship we're piloting throughout the event
// variants (i.e. old versions of the sub) will be properly named
// and placed into other files so that we maintain this as the
// current sub
type Submarine struct {
	depth    int
	position int
	aim      int
}

func NewSub() *Submarine {
	return &Submarine{depth: 0, position: 0, aim: 0}
}

//
// MOVEMENT FUNCS
//

func (s *Submarine) AimDown(deg int) {
	s.aim += deg
}

func (s *Submarine) AimUp(deg int) {
	s.aim -= deg
}

func (s *Submarine) Forward(dist int) {
	s.position += dist
	s.depth += s.aim * dist
}

//
// GETTERS
//

func (s *Submarine) GetDepth() int {
	return s.depth
}

func (s *Submarine) GetPosition() int {
	return s.position
}

func (s *Submarine) GetAim() int {
	return s.aim
}
