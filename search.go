package gosort

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
