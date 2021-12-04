package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func ReadFile(fname string) []string {
	file, err := os.Open(fname)
	Check(err)
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
		Check(err)
		vals = append(vals, num)
	}
	return vals
}

func LinesToStringInt(lines []string) ([]string, []int) {
	ss := make([]string, 0)
	is := make([]int, 0)

	for _, line := range lines {
		s, i := SingleLineToStringInt(line)
		ss = append(ss, s)
		is = append(is, i)
	}
	return ss, is
}

func SingleLineToStringInt(line string) (string, int) {
	splitLine := strings.Split(line, " ")

	s := splitLine[0]
	i, err := strconv.Atoi(splitLine[1])
	Check(err)

	return s, i
}

func BingoPullsStringToInt(pullStrings []string) []int {
	pulls := make([]int, 0)
	for _, pull := range pullStrings {
		p, err := strconv.Atoi(pull)
		Check(err)
		pulls = append(pulls, p)
	}
	return pulls
}
