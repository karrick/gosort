package gosort

// heapsort performs an in-place sort of the provided values.
func heapsort(values []int) {
	lv := len(values)

	// Outside loop is O(n), and we can skip first item.
	for i := 1; i < lv; i++ {

		// Inside is O(log n), bubbling up new value while it is smaller than
		// its parent value.
		var bubbleIndex = i
		newValue := values[bubbleIndex]

		for {
			parentIndex := (bubbleIndex - 1) >> 2
			parentValue := values[parentIndex]

			if parentValue < newValue {
				break // cannot bubble above a smaller value
			}

			values[bubbleIndex] = parentValue // move parent value down
			bubbleIndex = parentIndex         // next iteration look at parent position

			if bubbleIndex == 0 {
				break // already bubbled all the way up
			}
		}

		values[bubbleIndex] = newValue
	}
}
