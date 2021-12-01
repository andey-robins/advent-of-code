package main

import (
	"fmt"
	"log"

	"github.com/andey-robins/advent-of-code/dayOne"
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
	}
}
