package gosort

// heapsort performs an in-place sort of the provided values.
func heapsort(values []int) {
	lv := len(values)

	// STAGE 1: re-order elements to satisfy the heap property by creating a max
	// heap, where no child has a greater value than its parent.
	debug("stage 1: build max heap: %v\n", values)

	// Outside loop is O(n), and we can skip first item.
	for i := 1; i < lv; i++ {
		debug("\ti: %d\n", i)

		// Inside is O(log n), bubbling up new value while it is larger than its
		// parent.
		j := i
		v := values[j]
		debug("\tj: %v; value: %v\n", j, v)

		for j > 0 {
			parentIndex := (j - 1) >> 1
			parentValue := values[parentIndex]
			debug("\tparent index: %v\n\tparent value: %v\n", parentIndex, parentValue)

			if v < parentValue {
				debug("\tcannot bubble above a larger value\n")
				break
			}

			values[j] = parentValue // move parent value down
			j = parentIndex         // next iteration look at parent position
		}

		if j != i {
			values[j] = v
		}
	}
	// POST: values is a max-heap where no child has a value greater than its
	// parent.

	// STAGE 2: O(n) selection sort of largest element by walking heap to get
	// largest value.
	debug("stage 2: sort max heap: %v\n", values)

	for i := lv - 1; i > 0; i-- {

		// Swap first and final element.
		t := values[i]
		values[i] = values[0]

		// Re-establish max-heap property by sinking t until not greater than
		// parent.
		j := 0
		for {
			left := j<<1 + 1
			if left >= i {
				debug("\tleft >= i (%v >= %v)\n", left, i)
				break
			}
			vl := values[left]

			right := left + 1

			if right < i {
				if vr := values[right]; vl < vr {
					if vr < t {
						debug("\tvr < t (%v < %v)\n", vr, t)
						break
					}
					values[j] = vr
					j = right // go right
					continue
				}
			}

			if vl < t {
				debug("\tvl < t (%v < %v)\n", vl, t)
				break
			}
			values[j] = vl
			j = left // go left
		}

		values[j] = t
	}

	debug("complete: %v\n", values)
}
