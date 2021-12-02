package dayOne_test

import (
	"testing"

	"github.com/andey-robins/advent-of-code/sub"
	"github.com/andey-robins/advent-of-code/utils"
)

func TestPartOne(t *testing.T) {
	input := utils.ReadFile("./input.txt")
	readings := utils.LinesToInt(input)
	v := sub.CountDepthIncreases(readings, 1)
	if v != 1696 {
		t.FailNow()
	}
}

func TestPartTwo(t *testing.T) {
	input := utils.ReadFile("./input.txt")
	readings := utils.LinesToInt(input)
	v := sub.CountDepthIncreases(readings, 3)
	if v != 1737 {
		t.FailNow()
	}
}
