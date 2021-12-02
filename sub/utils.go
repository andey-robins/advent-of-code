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
