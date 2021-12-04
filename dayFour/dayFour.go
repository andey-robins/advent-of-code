package dayFour

import (
	"errors"
	"fmt"
	"strings"

	"github.com/andey-robins/advent-of-code/sub"
	"github.com/andey-robins/advent-of-code/utils"
)

func Run() {
	test()

	input := utils.ReadFile("./dayFour/input.txt")

	one, err := PartOne(input)
	utils.Check(err)
	two, err := PartTwo(input)
	utils.Check(err)

	fmt.Printf("Part One Solution: %v\n", one)
	fmt.Printf("Part Two Solution: %v\n", two)
}

func test() {
	input := utils.ReadFile("./dayFour/example.txt")

	// process pulls strings only once
	pullStrings := strings.Split(input[0], ",")
	pulls := utils.BingoPullsStringToInt(pullStrings)

	input = utils.ReadFile("./dayFour/example.txt")
	boards := sub.GenBingoBoardList(input[1:])

	winningBoard, lastPull, err := sub.BingoGetXthWinner(boards, pulls, 1)
	utils.Check(err)

	fmt.Printf("Test One: 4512 == %v\n", winningBoard.SumUncalled()*lastPull)

	// refresh for part two test
	input = utils.ReadFile("./dayFour/example.txt")
	boards = sub.GenBingoBoardList(input[1:])

	winningBoard, lastPull, err = sub.BingoGetXthWinner(boards, pulls, len(boards))
	utils.Check(err)

	fmt.Printf("Test Two: 1924 == %v\n", winningBoard.SumUncalled()*lastPull)
}

func PartOne(input []string) (int, error) {
	boards := sub.GenBingoBoardList(input[1:])
	pullStrings := strings.Split(input[0], ",")
	pulls := utils.BingoPullsStringToInt(pullStrings)

	winningBoard, lastPull, err := sub.BingoGetXthWinner(boards, pulls, 1)
	if err != nil {
		return 0, errors.New("no winning board found")
	}

	return winningBoard.SumUncalled() * lastPull, nil
}

func PartTwo(input []string) (int, error) {

	boards := sub.GenBingoBoardList(input[1:])
	pullStrings := strings.Split(input[0], ",")
	pulls := utils.BingoPullsStringToInt(pullStrings)

	winningBoard, lastPull, err := sub.BingoGetXthWinner(boards, pulls, len(boards))
	if err != nil {
		return 0, errors.New("no winning board found")
	}

	return winningBoard.SumUncalled() * lastPull, nil
}
