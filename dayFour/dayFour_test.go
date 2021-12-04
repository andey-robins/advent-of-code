package dayFour_test

import (
	"testing"

	"github.com/andey-robins/advent-of-code/dayFour"
	"github.com/andey-robins/advent-of-code/utils"
)

func TestPartOne(t *testing.T) {
	input := utils.ReadFile("./input.txt")
	v, err := dayFour.PartOne(input)
	if v != 38913 || err != nil {
		t.FailNow()
	}
}

func TestPartTwo(t *testing.T) {
	input := utils.ReadFile("./input.txt")
	v, err := dayFour.PartTwo(input)
	if v != 16836 || err != nil {
		t.FailNow()
	}
}
