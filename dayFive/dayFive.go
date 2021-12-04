package dayFive

import (
	"fmt"

	"github.com/andey-robins/advent-of-code/utils"
)

func Run() {
	testInput := utils.ReadFile("./dayX/example.txt")

	testOne := TestOne(testInput)
	testTwo := TestTwo(testInput)

	fmt.Printf("Test One Solution: XX == %v\n", testOne)
	fmt.Printf("Test Two Solution: XX == %v\n", testTwo)

	input := utils.ReadFile("./dayX/input.txt")

	one := PartOne(input)
	two := PartTwo(input)

	fmt.Printf("Part One Solution: %v\n", one)
	fmt.Printf("Part Two Solution: %v\n", two)
}

func TestOne(input []string) int {
	return 1
}

func TestTwo(input []string) int {
	return 2
}

func PartOne(input []string) int {
	return 1
}

func PartTwo(input []string) int {
	return 2
}
