package dayOne

import (
	"fmt"

	"github.com/andey-robins/advent-of-code/sub"
	"github.com/andey-robins/advent-of-code/utils"
)

func Run() {

	test()

	lines := utils.ReadFile("dayOne/input.txt")
	readings := utils.LinesToInt(lines)

	fmt.Printf("Part one answer: %v\n", sub.CountDepthIncreases(readings, 1))
	fmt.Printf("Part two answer: %v\n", sub.CountDepthIncreases(readings, 3))
}

func test() {
	lines := utils.ReadFile("dayOne/example.txt")
	readings := utils.LinesToInt(lines)

	fmt.Printf("Example test input yields: %v\n", sub.CountDepthIncreases(readings, 1))
	fmt.Printf("Example sliding window:    %v\n", sub.CountDepthIncreases(readings, 3))
}
