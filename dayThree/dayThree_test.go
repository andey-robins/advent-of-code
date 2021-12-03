package dayThree_test

import (
	"testing"

	"github.com/andey-robins/advent-of-code/dayThree"
	"github.com/andey-robins/advent-of-code/utils"
)

func TestPartOne(t *testing.T) {
	input := utils.ReadFile("./input.txt")
	v := dayThree.PartOne(input)
	if v != 3148794 {
		t.FailNow()
	}
}

func TestPartTwo(t *testing.T) {
	input := utils.ReadFile("./input.txt")
	v := dayThree.PartTwo(input)
	if v != 2795310 {
		t.FailNow()
	}
}
