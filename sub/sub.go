package sub

import (
	"strconv"

	"github.com/andey-robins/advent-of-code/utils"
)

// Submarine is the ship we're piloting throughout the event
// variants (i.e. old versions of the sub) will be properly named
// and placed into other files so that we maintain this as the
// current sub
type Submarine struct {
	depth    int
	position int
	aim      int
	gamma    int
	epsilon  int
	oxRating int
	coRating int
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
// COMPUTER ERROR CODE PROCESSING
//

func (s *Submarine) CalculateGamma(codes []string) {
	size := len(codes[0])
	counters := getErrorCodeCounts(codes, size)

	gamma := ""

	for _, x := range counters {
		if x > 0 {
			gamma += "1"
		} else {
			gamma += "0"
		}
	}

	g, err := strconv.ParseInt(gamma, 2, 64)
	utils.Check(err)
	s.gamma = int(g)
}

func (s *Submarine) CalculateEpsilon(codes []string) {
	size := len(codes[0])
	counters := getErrorCodeCounts(codes, size)

	epsilon := ""

	for _, x := range counters {
		if x > 0 {
			epsilon += "0"
		} else {
			epsilon += "1"
		}
	}

	e, err := strconv.ParseInt(epsilon, 2, 64)
	utils.Check(err)
	s.epsilon = int(e)
}

func (s *Submarine) CalculateOxRating(codes []string) {
	size := len(codes[0])
	counters := getErrorCodeCounts(codes, size)

	for idx := 0; len(codes) > 1; idx++ {
		codes = removeAtPos(codes, (counters[idx] < 0), idx)
		counters = getErrorCodeCounts(codes, size)
	}

	o, err := strconv.ParseInt(codes[0], 2, 64)
	utils.Check(err)
	s.oxRating = int(o)
}

func (s *Submarine) CalculateCoRating(codes []string) {
	size := len(codes[0])
	counters := getErrorCodeCounts(codes, size)

	for idx := 0; len(codes) > 1; idx++ {
		codes = removeAtPos(codes, (counters[idx] >= 0), idx)
		counters = getErrorCodeCounts(codes, size)
	}

	c, err := strconv.ParseInt(codes[0], 2, 64)
	utils.Check(err)
	s.coRating = int(c)
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

func (s *Submarine) GetGamma() int {
	return s.gamma
}

func (s *Submarine) GetEpsilon() int {
	return s.epsilon
}

func (s *Submarine) GetOxRating() int {
	return s.oxRating
}

func (s *Submarine) GetCoRating() int {
	return s.coRating
}
