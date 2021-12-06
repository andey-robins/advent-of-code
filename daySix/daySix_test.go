package daySix_test

import (
	"math/big"
	"testing"

	"github.com/andey-robins/advent-of-code/daySix"
	"github.com/andey-robins/advent-of-code/utils"
)

func TestTestOne(t *testing.T) {
	input := utils.ReadFile("./example.txt")
	v := daySix.PartOne(input)
	if v.Cmp(big.NewInt(5934)) != 0 {
		t.FailNow()
	}
}

func TestTestTwo(t *testing.T) {
	input := utils.ReadFile("./example.txt")
	v := daySix.PartTwo(input)
	if v.Cmp(big.NewInt(26984457539)) != 0 {
		t.FailNow()
	}
}

func TestPartOne(t *testing.T) {
	input := utils.ReadFile("./input.txt")
	v := daySix.PartOne(input)
	if v.Cmp(big.NewInt(379114)) != 0 {
		t.FailNow()
	}
}

func TestPartTwo(t *testing.T) {
	input := utils.ReadFile("./input.txt")
	v := daySix.PartTwo(input)
	if v.Cmp(big.NewInt(1702631502303)) != 0 {
		t.FailNow()
	}
}
