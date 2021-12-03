package dayX_test

import (
	"testing"

	"github.com/andey-robins/advent-of-code/utils"
)

func TestPartOne(t *testing.T) {
	input := utils.ReadFile("./input.txt")
	v := dayX.PartOne(input)
	if v != 1111 {
		t.FailNow()
	}
}

func TestPartTwo(t *testing.T) {
	input := utils.ReadFile("./input.txt")
	v := dayX.PartTwo(input)
	if v != 2222 {
		t.FailNow()
	}
}
