package gosort

// merge sorts list a, creating a scratch list as a buffer.
func merge(a []int) {
	n := len(a)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = a[i]
	}
	mergeSplit(b, a, 0, n)
}

// mergeSplit recursively splits b into sub lists until they are no more than a
// single element wide, and merges the results into a.
func mergeSplit(b, a []int, begin, end int) {
	if end-begin < 2 {
		return // single item is already sorted
	}

	middle := (begin + end) >> 1
	mergeSplit(a, b, begin, middle)
	mergeSplit(a, b, middle, end)
	mergeLists(b, a, begin, middle, end)
}

// mergeLists merges two sublists from a[begin:middle-1] and a[middle:end-1]
// into b.
func mergeLists(a, b []int, begin, middle, end int) {
	i := begin
	j := middle

	for k := begin; k < end; k++ {
		if i < middle && (j >= end || a[i] < a[j]) {
			b[k] = a[i]
			i++
		} else {
			b[k] = a[j]
			j++
		}
	}
}
