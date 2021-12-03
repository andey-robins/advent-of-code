package dayFour

import (
	"fmt"

	"github.com/andey-robins/advent-of-code/utils"
)

func Run() {
	test()

	input := utils.ReadFile("./dayFour/input.txt")

	one := PartOne(input)
	two := PartTwo(input)

	fmt.Printf("Part One Solution: %v\n", one)
	fmt.Printf("Part Two Solution: %v\n", two)
}

func test() {

}

func PartOne(input []string) int {
	return 1
}

func PartTwo(input []string) int {
	return 2
}
