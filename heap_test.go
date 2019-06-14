package gosort

import (
	"math/rand"
	"testing"
)

func TestHeap(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		var values []int
		heapsort(values)
		ensureIntSlicesMatch(t, values, nil)
	})

	t.Run("one", func(t *testing.T) {
		values := []int{13}
		heapsort(values)
		ensureIntSlicesMatch(t, values, []int{13})
	})

	t.Run("two", func(t *testing.T) {
		t.Run("ordered", func(t *testing.T) {
			values := []int{13, 42}
			heapsort(values)
			ensureIntSlicesMatch(t, values, []int{13, 42})
		})

		t.Run("unordered", func(t *testing.T) {
			values := []int{42, 13}
			heapsort(values)
			ensureIntSlicesMatch(t, values, []int{13, 42})
		})
	})

	t.Run("three", func(t *testing.T) {
		t.Run("ordered", func(t *testing.T) {
			values := []int{13, 42, 97}
			heapsort(values)
			ensureIntSlicesMatch(t, values, []int{13, 42, 97})
		})

		t.Run("unordered", func(t *testing.T) {
			values := []int{97, 42, 13}
			heapsort(values)
			ensureIntSlicesMatch(t, values, []int{13, 42, 97})
		})
	})

	t.Run("randomized", func(t *testing.T) {
		const max = 1000

		want := make([]int, max)
		for i := 0; i < max; i++ {
			want[i] = i
		}

		values := rand.Perm(max) // generate a randomized list
		heapsort(values)

		ensureIntSlicesMatch(t, values, want)
	})
}
