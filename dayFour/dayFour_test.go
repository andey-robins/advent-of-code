package dayFour_test

import (
	"testing"

	"github.com/andey-robins/advent-of-code/dayFour"
	"github.com/andey-robins/advent-of-code/utils"
)

func TestPartOne(t *testing.T) {
	input := utils.ReadFile("./input.txt")
	v := dayFour.PartOne(input)
	if v != 1111 {
		t.FailNow()
	}
}

func TestPartTwo(t *testing.T) {
	input := utils.ReadFile("./input.txt")
	v := dayFour.PartTwo(input)
	if v != 2222 {
		t.FailNow()
	}
}
