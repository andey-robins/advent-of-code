package dayFive_test

import (
	"testing"

	"github.com/andey-robins/advent-of-code/dayFive"
	"github.com/andey-robins/advent-of-code/utils"
)

func TestTestOne(t *testing.T) {
	input := utils.ReadFile("./input.txt")
	v := dayFive.TestOne(input)
	if v != 1 {
		t.FailNow()
	}
}

func TestTestTwo(t *testing.T) {
	input := utils.ReadFile("./input.txt")
	v := dayFive.TestTwo(input)
	if v != 2 {
		t.FailNow()
	}
}

func TestPartOne(t *testing.T) {
	input := utils.ReadFile("./input.txt")
	v := dayFive.PartOne(input)
	if v != 1 {
		t.FailNow()
	}
}

func TestPartTwo(t *testing.T) {
	input := utils.ReadFile("./input.txt")
	v := dayFive.PartTwo(input)
	if v != 2 {
		t.FailNow()
	}
}
