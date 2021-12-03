package sub

// count the number of increases in depth of sonar readings
// allows for sliding window size defined by the arg window
func CountDepthIncreases(sonar []int, window int) (count int) {
	count = 0
	last := sumWindow(sonar, 0, window)

	for i := 1; i < len(sonar)-window+1; i++ {
		currWindow := sumWindow(sonar, i, window)
		if currWindow > last {
			count++
		}
		last = currWindow
	}

	return count
}

// given a set of sonar readings, sum the window of size count starting
// at location idx
func sumWindow(sonar []int, idx int, count int) int {
	sum := 0
	for i := idx; i < idx+count; i++ {
		sum += sonar[i]
	}
	return sum
}

// given a list of error codes, and a decision to remove codes with a zero
// or one at position pos, return a new set of codes following the given
// ruleset
func removeAtPos(lines []string, removeOnes bool, pos int) []string {
	res := make([]string, 0)

	for _, line := range lines {
		if removeOnes {
			if line[pos] == '0' {
				res = append(res, line)
			}
		} else {
			if line[pos] == '1' {
				res = append(res, line)
			}
		}
	}

	return res
}

// return the count of the number of 1s and 0s at each position for
// a set of error codes. if the number at index i is positive, there
// are more 1s than 0s. if it's negative, then the opposite
func getErrorCodeCounts(input []string, size int) []int {
	counters := make([]int, size)

	for _, line := range input {
		for i, n := range line {
			if n == '1' {
				counters[i]++
			} else {
				counters[i]--
			}
		}
	}
	return counters
}
