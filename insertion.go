package gosort

func insertionsort1(values []int) {
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

func insertionsort2(values []int) {
	lv := len(values)

	for i := 1; i < lv; i++ {
		// Insert the i'th value into the sorted list from index 0 to i-1.
		v := values[i]

		if !(v < values[i-1]) {
			continue
		}

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

func searchLeftMostGreaterThan(v int, values []int) int {
	lo := 0
	hi := len(values) - 1
	for lo <= hi {
		m := (lo + hi) >> 1
		if v < values[m] {
			hi = m - 1
		} else {
			lo = m + 1
		}
	}
	return lo
}

func insertionsort3(values []int) {
	lv := len(values)

	for i := 1; i < lv; i++ {
		// Insert the i'th value into the sorted list from index 0 to i-1.
		v := values[i]

		hi := i - 1

		if !(v < values[hi]) {
			continue
		}

		// lo := searchLeftMostGreaterThan(v, values[:i-1])
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
