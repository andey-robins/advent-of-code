package dayFive_test

import (
	"testing"

	"github.com/andey-robins/advent-of-code/dayFive"
	"github.com/andey-robins/advent-of-code/utils"
)

func TestTestOne(t *testing.T) {
	input := utils.ReadFile("./example.txt")
	v := dayFive.PartOne(input)
	if v != 5 {
		t.FailNow()
	}
}

func TestTestTwo(t *testing.T) {
	input := utils.ReadFile("./example.txt")
	v := dayFive.PartTwo(input)
	if v != 12 {
		t.FailNow()
	}
}

func TestPartOne(t *testing.T) {
	input := utils.ReadFile("./input.txt")
	v := dayFive.PartOne(input)
	if v != 6461 {
		t.FailNow()
	}
}

func TestPartTwo(t *testing.T) {
	input := utils.ReadFile("./input.txt")
	v := dayFive.PartTwo(input)
	if v != 18065 {
		t.FailNow()
	}
}
