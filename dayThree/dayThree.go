package dayThree

import (
	"fmt"

	"github.com/andey-robins/advent-of-code/sub"
	"github.com/andey-robins/advent-of-code/utils"
)

func Run() {
	test()

	input := utils.ReadFile("./dayThree/input.txt")

	one := PartOne(input)
	two := PartTwo(input)

	fmt.Printf("Part One Solution: %v\n", one)
	fmt.Printf("Part Two Solution: %v\n", two)
}

func test() {
	input := utils.ReadFile("./dayThree/example.txt")

	sub := sub.NewSub()

	sub.CalculateCoRating(input)
	sub.CalculateEpsilon(input)
	sub.CalculateGamma(input)
	sub.CalculateOxRating(input)

	fmt.Printf("Gamma Value:         22 == %v\n", sub.GetGamma())
	fmt.Printf("Epsilon Value:       9  == %v\n", sub.GetEpsilon())
	fmt.Printf("O2 Generator Rating: 23 == %v\n", sub.GetOxRating())
	fmt.Printf("CO2 Scrubber Rating: 10 == %v\n", sub.GetCoRating())
}

func PartOne(input []string) int {

	sub := sub.NewSub()

	sub.CalculateEpsilon(input)
	sub.CalculateGamma(input)

	return sub.GetEpsilon() * sub.GetGamma()
}

func PartTwo(input []string) int {

	sub := sub.NewSub()

	sub.CalculateOxRating(input)
	sub.CalculateCoRating(input)

	return sub.GetOxRating() * sub.GetCoRating()
}
