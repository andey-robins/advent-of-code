package dayOne

import (
	"fmt"
	"log"
)

func Run() {
	var i int
	fmt.Println("Enter 1 for part one and anything else for part 2")
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("\n\n")

	if i == 1 {
		partOne()
	} else {
		partTwo()
	}
}

func partOne() {
	fmt.Println("Part 1")
}

func partTwo() {
	fmt.Println("Part 2")
}
