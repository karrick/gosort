package gosort

func insertionsort1(values []int) {
	debug("insertion: %v\n", values)
	lv := len(values)

	for i := 1; i < lv; i++ {
		// Insert the i'th value into the sorted list from index 0 to i-1.
		v := values[i]
		debug("insert %v into sorted list: %v\n", v, values[:i])

		var j int

		for j = i - 1; j >= 0; j-- {
			debug("j: %v\n", j)
			if v < values[j] {
				debug("before swap: %v\n", values)
				values[j+1] = values[j]
				debug("after  swap: %v\n", values)
				continue
			}
			debug("v is greater than or equal: %v\n", v)
			break
		}

		j++
		debug("post j: %d\n", j)
		values[j] = v
	}

	debug("done: %v\n", values)
}

func insertionsort2(values []int) {
	debug("insertion: %v\n", values)
	lv := len(values)

	for i := 1; i < lv; i++ {
		// Insert the i'th value into the sorted list from index 0 to i-1.
		v := values[i]
		debug("insert %v into sorted list: %v\n", v, values[:i])

		if !(v < values[i-1]) {
			debug("shortcut: %v\n", v)
			continue
		}

		var j int

		for j = i - 1; j >= 0; j-- {
			debug("j: %v\n", j)
			if v < values[j] {
				debug("before swap: %v\n", values)
				values[j+1] = values[j]
				debug("after  swap: %v\n", values)
				continue
			}
			debug("v is greater than or equal: %v\n", v)
			break
		}

		j++
		debug("post j: %d\n", j)
		values[j] = v
	}

	debug("done: %v\n", values)
}

func insertionsort3(values []int) {
	debug("insertion: %v\n", values)
	lv := len(values)

	for i := 1; i < lv; i++ {
		// Insert the i'th value into the sorted list from index 0 to i-1.
		v := values[i]
		debug("\tinsert %v into sorted list: %v; tail: %v\n", v, values[:i], values[i:])

		if !(v < values[i-1]) {
			debug("\t\tshortcut: %v\n", v)
			continue
		}

		lo := 0
		hi := i - 1
		debug("\t\tbinary search from %v to %v to find the insertion position\n", lo, hi)

		for lo <= hi {
			m := (lo + hi) >> 1
			vm := values[m]
			debug("\t\t\tlo: %v; hi: %v; m: %v; v: %v; vm: %v\n", lo, hi, m, v, vm)
			if v < vm {
				hi = m - 1
			} else if vm < v {
				lo = m + 1
			} else {
				break
			}
		}
		// POST: lo is index of insertion

		for !(v < values[lo]) {
			debug("\tequal: walking forward until v < vm\n")
			lo++
		}

		debug("lo: %v; values: %v\n", lo, values)
		copy(values[lo+1:i+1], values[lo:i+1])
		values[lo] = v
	}

	debug("done: %v\n", values)
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

func insertionsort4(values []int) {
	debug("insertion: %v\n", values)
	lv := len(values)

	for i := 1; i < lv; i++ {
		// Insert the i'th value into the sorted list from index 0 to i-1.
		v := values[i]
		debug("\tinsert %v into sorted list: %v; tail: %v\n", v, values[:i], values[i:])

		hi := i - 1

		if !(v < values[hi]) {
			debug("\t\tshortcut: %v\n", v)
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

		debug("lo: %v; values: %v\n", lo, values)
		copy(values[lo+1:i+1], values[lo:i+1])
		values[lo] = v
	}

	debug("done: %v\n", values)
}
