package daySeven

import (
	"fmt"

	"github.com/andey-robins/advent-of-code/utils"
)

func Run() {
	testInput := utils.ReadFile("./daySeven/example.txt")

	testOne := PartOne(testInput)
	fmt.Printf("Test One Solution: XX == %v\n", testOne)
	testTwo := PartTwo(testInput)
	fmt.Printf("Test Two Solution: XX == %v\n", testTwo)

	input := utils.ReadFile("./daySeven/input.txt")

	one := PartOne(input)
	fmt.Printf("Part One Solution: %v\n", one)
	two := PartTwo(input)
	fmt.Printf("Part Two Solution: %v\n", two)
}

func PartOne(input []string) int {

	// calc median
	// sum abs value of median-val

	return 1
}

func PartTwo(input []string) int {

	// calc average, round down
	// sum the sum of values up to abs avg-val

	return 2
}
