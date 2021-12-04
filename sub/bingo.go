package sub

import (
	"errors"
	"strconv"
	"strings"

	"github.com/andey-robins/advent-of-code/utils"
)

type Board struct {
	nums [][]Spot
}

type Spot struct {
	num    int
	called bool
}

func NewBingoBoard() *Board {
	tempBoard := &Board{}
	for i := 0; i < 5; i++ {
		tempBoard.nums = make([][]Spot, 5)
	}
	return tempBoard
}

func GenBingoBoardList(input []string) []*Board {
	boards := make([]*Board, 0)
	// fill in boards with all valid boards
	tempBoard := NewBingoBoard()
	for _, line := range input {
		if line == "" {
			boards = append(boards, tempBoard)
			tempBoard = NewBingoBoard()
		} else {
			tempBoard.AddLine(line)
		}
	}
	boards = append(boards, tempBoard)
	boards = boards[1:]

	return boards
}

func (b *Board) Call(n int) {
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if b.nums[x][y].num == n {
				b.nums[x][y].called = true
			}
		}
	}
}

func (b *Board) CheckWin() bool {
	for _, row := range b.nums {
		if listIsWin(row) {
			return true
		}
	}

	for x := 0; x < 5; x++ {
		col := make([]Spot, 0)
		for y := 0; y < 5; y++ {
			col = append(col, b.nums[y][x])
		}
		if listIsWin(col) {
			return true
		}
	}
	return false
}

func (b *Board) AddLine(line string) {
	numsStrings := strings.Split(line, " ")
	idx := 0
	for _, n := range numsStrings {
		if n != "" {
			x, err := strconv.Atoi(n)
			utils.Check(err)
			b.nums[idx] = append(b.nums[idx], Spot{num: x, called: false})
			idx++
		}
	}
}

func (b *Board) SumUncalled() int {
	sum := 0
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if !b.nums[x][y].called {
				sum += b.nums[x][y].num
			}
		}
	}
	return sum
}

func BingoGetXthWinner(boards []*Board, pulls []int, x int) (*Board, int, error) {
	startingSize := len(boards)
	for _, pull := range pulls {
		boardList := boards
		for _, board := range boardList {
			board.Call(pull)
		}
		for i := len(boardList) - 1; i >= 0; i-- {
			if boardList[i].CheckWin() {
				lastRemoved := boardList[i]
				boards = BingoBoardRemove(boards, i)
				if len(boards) == startingSize-x {
					return lastRemoved, pull, nil
				}
			}
		}
	}
	return nil, 0, errors.New("xth winning board not found")
}

func BingoBoardRemove(slice []*Board, s int) []*Board {
	return append(slice[:s], slice[s+1:]...)
}

func listIsWin(list []Spot) bool {
	for _, i := range list {
		if !i.called {
			return false
		}
	}
	return true
}
