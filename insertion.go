package gosort

// insertionsort1 is traditional insertion sort with no optimizations.
func insertionsort1(values []int) {
	lv := len(values)

	for i := 1; i < lv; i++ {
		// Insert the i'th value into the sorted list from index 0 to i-1.
		v := values[i]

		for j := i - 1; j >= 0; j-- {
			if v < values[j] {
				values[j+1] = values[j]
				values[j] = v
				continue
			}
			break
		}
	}
}

// insertionsort2 is slight optimization in that it only copies value into place
// after search for target location has been found, rather than doing as many
// assignments.
func insertionsort2(values []int) {
	lv := len(values)

	for i := 1; i < lv; i++ {
		// Insert the i'th value into the sorted list from index 0 to i-1.
		v := values[i]

		var j int

		for j = i - 1; j >= 0; j-- {
			if v < values[j] {
				values[j+1] = values[j]
				continue
			}
			break
		}

		j++
		values[j] = v
	}
}

// insertionsort3 provides two optimizations.  First, when new value is greater
// than right most value in sorted partition, this merely appends the value to
// the sorted partition.  The second optimization is that when the new value is
// not greater than the right most value in the sorted partition, it uses a
// binary search to find the left most value that is greater than the new value
// to locate the insertion position, copies all values over one slot to the
// right, then places the new value into position.
func insertionsort3(values []int) {
	lv := len(values)

	for i := 1; i < lv; i++ {
		// Insert the i'th value into the sorted list from index 0 to i-1.
		v := values[i]

		hi := i - 1

		if !(v < values[hi]) {
			continue // optimization for pre-sorted lists
		}

		// lo := searchLeftMostGreaterThan(v, values[:hi)
		lo := 0
		for lo <= hi {
			m := (lo + hi) >> 1
			if v < values[m] {
				hi = m - 1
			} else {
				lo = m + 1
			}
		}
		// return lo

		copy(values[lo+1:i+1], values[lo:i+1])
		values[lo] = v
	}
}
