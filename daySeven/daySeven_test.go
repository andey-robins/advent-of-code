package daySeven_test

import (
	"testing"

	"github.com/andey-robins/advent-of-code/daySeven"
	"github.com/andey-robins/advent-of-code/utils"
)

func TestTestOne(t *testing.T) {
	input := utils.ReadFile("./example.txt")
	v := daySeven.PartOne(input)
	if v != 1 {
		t.FailNow()
	}
}

func TestTestTwo(t *testing.T) {
	input := utils.ReadFile("./example.txt")
	v := daySeven.PartTwo(input)
	if v != 2 {
		t.FailNow()
	}
}

func TestPartOne(t *testing.T) {
	input := utils.ReadFile("./input.txt")
	v := daySeven.PartOne(input)
	if v != 1 {
		t.FailNow()
	}
}

func TestPartTwo(t *testing.T) {
	input := utils.ReadFile("./input.txt")
	v := daySeven.PartTwo(input)
	if v != 2 {
		t.FailNow()
	}
}
