package gosort

// heapsort performs an in-place sort of the provided values.
func heapsort(values []int) {
	lv := len(values)

	// STAGE 1: Re-order elements to satisfy the heap property by creating a max
	// heap, where no child has a greater value than its parent.
	debug("stage 1: build max heap: %v\n", values)

	// Outside loop is O(n), and we can skip first item because it is already
	// sorted when in a list by itself.
	for i := 1; i < lv; i++ {
		debug("\ti: %d\n", i)

		// Inside loop is O(log n), bubbling up new value while it is larger
		// than its parent, but since these elements are sorted in a heap, it
		// only needs to bubble up a max of log n times rather than n.
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

	// STAGE 2: The first element in the max heap is the largest item in the
	// list.  Iteratively take the first element and place it at the end of the
	// list, re-balance the heap, and then shrink the list by one element.
	debug("stage 2: sort max heap: %v\n", values)

	// Outside loop is O(n), where we take the first and largest element in the
	// max heap, place it at the end of the list, then use a loop to re-balance
	// the list.
	for i := lv - 1; i > 0; i-- {

		// Swap first and final element.
		t := values[i]        // save final element
		values[i] = values[0] // move largest element into final position

		// Find a new home for the saved final element by walking down the max
		// heap until it is smaller than some other element.
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
