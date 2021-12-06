package daySix

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/andey-robins/advent-of-code/utils"
)

type Generation struct {
	fish []big.Int
}

func NewGeneration(nums []int) *Generation {
	fish := make([]big.Int, 9)

	for i := 0; i < 9; i++ {
		fish[i] = *big.NewInt(0)
	}

	for _, n := range nums {
		newCount := big.NewInt(0)
		newCount.Add(&fish[n], big.NewInt(1))
		fish[n] = *newCount
	}

	return &Generation{fish}
}

func (g *Generation) calcNextGen() {
	splitting := g.fish[0]

	for i := 1; i < 9; i++ {
		g.fish[i-1] = g.fish[i]
	}

	newSix := big.NewInt(0)
	newSix.Add(&g.fish[6], &splitting)
	g.fish[6] = *newSix
	g.fish[8] = splitting
}

func (g *Generation) Simulate(len int) {
	for i := 0; i < len; i++ {
		g.calcNextGen()
	}
}

func (g *Generation) Sum() *big.Int {
	sum := big.NewInt(0)
	for i := 0; i < 9; i++ {
		sum.Add(sum, &g.fish[i])
	}
	return sum
}

func Run() {
	testInput := utils.ReadFile("./daySix/example.txt")

	testOne := PartOne(testInput)
	fmt.Printf("Test One Solution: 5934 == %v\n", testOne)
	testTwo := PartTwo(testInput)
	fmt.Printf("Test Two Solution: 26984457539 == %v\n", testTwo)

	input := utils.ReadFile("./daySix/input.txt")

	one := PartOne(input)
	fmt.Printf("Part One Solution: %v\n", one)
	two := PartTwo(input)
	fmt.Printf("Part Two Solution: %v\n", two)
}

func PartOne(input []string) *big.Int {
	startingAges := getStartingAges(input[0])
	sim := NewGeneration(startingAges)
	sim.Simulate(80)
	return sim.Sum()
}

func PartTwo(input []string) *big.Int {
	startingAges := getStartingAges(input[0])
	sim := NewGeneration(startingAges)
	sim.Simulate(256)
	return sim.Sum()
}

func getStartingAges(s string) []int {
	fishStrings := strings.Split(s, ",")
	startingAges := make([]int, 0)

	for _, fishString := range fishStrings {
		if x, err := strconv.Atoi(fishString); err != nil {
			panic(err)
		} else {
			startingAges = append(startingAges, x)
		}
	}

	return startingAges
}
