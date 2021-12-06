package main

import (
	"fmt"
	"log"

	"github.com/andey-robins/advent-of-code/dayFive"
	"github.com/andey-robins/advent-of-code/dayFour"
	"github.com/andey-robins/advent-of-code/dayOne"
	"github.com/andey-robins/advent-of-code/daySix"
	"github.com/andey-robins/advent-of-code/dayThree"
	"github.com/andey-robins/advent-of-code/dayTwo"
)

func main() {
	var i int
	fmt.Print("Enter Day to Run: ")
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		log.Fatalln(err)
	}

	switch i {
	case 1:
		dayOne.Run()
	case 2:
		dayTwo.Run()
	case 3:
		dayThree.Run()
	case 4:
		dayFour.Run()
	case 5:
		dayFive.Run()
	case 6:
		daySix.Run()
	}
}
