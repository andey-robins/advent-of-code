package dayTwo_test

import (
	"testing"

	"github.com/andey-robins/advent-of-code/dayTwo"
	"github.com/andey-robins/advent-of-code/utils"
)

func TestPartOne(t *testing.T) {
	input := utils.ReadFile("./input.txt")
	v := dayTwo.PartOne(input)
	if v != 1947824 {
		t.FailNow()
	}
}

func TestPartTwo(t *testing.T) {
	input := utils.ReadFile("./input.txt")
	v := dayTwo.PartTwo(input)
	if v != 1813062561 {
		t.FailNow()
	}
}
