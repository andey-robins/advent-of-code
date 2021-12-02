package dayTwo

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/andey-robins/advent-of-code/sub"
	"github.com/andey-robins/advent-of-code/utils"
)

func Run() {
	test()

	input := utils.ReadFile("./dayTwo/input.txt")

	one := PartOne(input)
	two := PartTwo(input)

	fmt.Printf("Part One Answer: %v\n", one)
	fmt.Printf("Part Two Answer: %v\n", two)
}

func test() {
	input := utils.ReadFile("./dayTwo/example.txt")

	testSimpleSub := sub.NewSimpleSub()
	testSub := sub.NewSub()

	for _, line := range input {
		splitLine := strings.Split(line, " ")

		move := splitLine[0]
		dist, err := strconv.Atoi(splitLine[1])
		utils.Check(err)

		switch move {
		case "up":
			testSimpleSub.Rise(dist)
			testSub.AimUp(dist)
		case "down":
			testSimpleSub.Dive(dist)
			testSub.AimDown(dist)
		case "forward":
			testSimpleSub.Forward(dist)
			testSub.Forward(dist)
		default:
			fmt.Printf("Unable to read: %v", move)
		}
	}

	fmt.Printf("Example simple sub result: 150 == %v\n", testSimpleSub.GetDepth()*testSimpleSub.GetPosition())
	fmt.Printf("Example sub result       : 900 == %v\n", testSub.GetDepth()*testSub.GetPosition())
}

func PartOne(input []string) int {

	sub := sub.NewSimpleSub()

	for _, line := range input {
		move, dist := utils.SingleLineToStringInt(line)

		switch move {
		case "up":
			sub.Rise(dist)
		case "down":
			sub.Dive(dist)
		case "forward":
			sub.Forward(dist)
		default:
			fmt.Printf("Unable to read: %v", move)
		}
	}

	return sub.GetDepth() * sub.GetPosition()
}

func PartTwo(input []string) int {

	sub := sub.NewSub()

	for _, line := range input {
		move, dist := utils.SingleLineToStringInt(line)

		switch move {
		case "up":
			sub.AimUp(dist)
		case "down":
			sub.AimDown(dist)
		case "forward":
			sub.Forward(dist)
		default:
			fmt.Printf("Unable to read: %v", move)
		}
	}

	return sub.GetDepth() * sub.GetPosition()
}
