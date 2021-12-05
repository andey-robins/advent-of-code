package dayFive

import (
	"fmt"
	"math"

	"github.com/andey-robins/advent-of-code/utils"
)

func Run() {
	testInput := utils.ReadFile("./dayFive/example.txt")

	testOne := PartOne(testInput)
	fmt.Printf("Test One Solution: 5  == %v\n", testOne)
	testTwo := PartTwo(testInput)
	fmt.Printf("Test Two Solution: 12 == %v\n", testTwo)

	input := utils.ReadFile("./dayFive/input.txt")

	one := PartOne(input)
	fmt.Printf("Part One Solution: %v\n", one)
	two := PartTwo(input)
	fmt.Printf("Part Two Solution: %v\n", two)
}

func PartOne(input []string) int {
	lines := utils.VentsReadCoords(input)
	size := utils.GetMaxSize(lines) + 1

	oceanFloor := make([][]int, size)
	for i := 0; i < size; i++ {
		oceanFloor[i] = make([]int, size)
	}

	for _, line := range lines {
		oceanFloor = drawLine(line, oceanFloor, true)
	}

	return countHits(oceanFloor)
}

func PartTwo(input []string) int {
	lines := utils.VentsReadCoords(input)
	size := utils.GetMaxSize(lines) + 1

	oceanFloor := make([][]int, size)
	for i := 0; i < size; i++ {
		oceanFloor[i] = make([]int, size)
	}

	for _, line := range lines {
		oceanFloor = drawLine(line, oceanFloor, false)
	}

	return countHits(oceanFloor)
}

func countHits(floor [][]int) int {
	hits := 0
	for _, row := range floor {
		for _, elem := range row {
			if elem >= 2 {
				hits++
			}
		}
	}
	return hits
}

func drawLine(line utils.Line, floor [][]int, isPartOne bool) [][]int {
	if line.Xs[0] == line.Xs[1] {
		var start int
		var end int

		if line.Ys[0] < line.Ys[1] {
			start = line.Ys[0]
			end = line.Ys[1]
		} else {
			start = line.Ys[1]
			end = line.Ys[0]
		}

		// draw horizontal
		for i := start; i <= end; i++ {
			floor[i][line.Xs[0]]++
		}
	} else if line.Ys[0] == line.Ys[1] {
		var start int
		var end int

		if line.Xs[0] < line.Xs[1] {
			start = line.Xs[0]
			end = line.Xs[1]
		} else {
			start = line.Xs[1]
			end = line.Xs[0]
		}

		// draw horizontal
		for i := start; i <= end; i++ {
			floor[line.Ys[0]][i]++
		}
	} else if !isPartOne {
		var slope int
		// draw diagonals
		if line.Ys[0] < line.Ys[1] {
			if line.Xs[0] < line.Xs[1] {
				slope = 1
			} else {
				slope = -1
			}
		} else {
			if line.Xs[0] < line.Xs[1] {
				slope = -1
			} else {
				slope = 1
			}
		}

		startX := int(math.Min(float64(line.Xs[0]), float64(line.Xs[1])))
		endX := int(math.Max(float64(line.Xs[0]), float64(line.Xs[1])))

		var y int
		if startX == line.Xs[0] {
			y = line.Ys[0]
		} else {
			y = line.Ys[1]
		}
		for x := startX; x <= endX; x++ {
			floor[y][x]++
			y += slope
		}
	}
	return floor
}
