package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func ReadFile(fname string) []string {
	file, err := os.Open(fname)
	check(err)
	defer file.Close()

	contents := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		panic(err)
	}

	return contents
}

func LinesToInt(lines []string) []int {
	vals := make([]int, 0)
	for _, line := range lines {
		num, err := strconv.Atoi(line)
		check(err)
		vals = append(vals, num)
	}
	return vals
}
