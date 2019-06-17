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

	t.Run("a", func(t *testing.T) {
		values := []int{13}
		heapsort(values)
		ensureIntSlicesMatch(t, values, []int{13})
	})

	t.Run("two", func(t *testing.T) {
		t.Run("ab", func(t *testing.T) {
			values := []int{13, 42}
			heapsort(values)
			ensureIntSlicesMatch(t, values, []int{13, 42})
		})
		t.Run("ba", func(t *testing.T) {
			values := []int{42, 13}
			heapsort(values)
			ensureIntSlicesMatch(t, values, []int{13, 42})
		})
	})

	t.Run("three", func(t *testing.T) {
		t.Run("abc", func(t *testing.T) {
			values := []int{13, 42, 97}
			heapsort(values)
			ensureIntSlicesMatch(t, values, []int{13, 42, 97})
		})
		t.Run("bca", func(t *testing.T) {
			values := []int{42, 97, 13}
			heapsort(values)
			ensureIntSlicesMatch(t, values, []int{13, 42, 97})
		})
		t.Run("cab", func(t *testing.T) {
			values := []int{97, 13, 42}
			heapsort(values)
			ensureIntSlicesMatch(t, values, []int{13, 42, 97})
		})
		t.Run("acb", func(t *testing.T) {
			values := []int{13, 97, 42}
			heapsort(values)
			ensureIntSlicesMatch(t, values, []int{13, 42, 97})
		})
		t.Run("cba", func(t *testing.T) {
			values := []int{97, 42, 13}
			heapsort(values)
			ensureIntSlicesMatch(t, values, []int{13, 42, 97})
		})
		t.Run("bac", func(t *testing.T) {
			values := []int{42, 13, 97}
			heapsort(values)
			ensureIntSlicesMatch(t, values, []int{13, 42, 97})
		})
		t.Run("acb", func(t *testing.T) {
			values := []int{13, 97, 42}
			heapsort(values)
			ensureIntSlicesMatch(t, values, []int{13, 42, 97})
		})
	})

	t.Run("ten", func(t *testing.T) {
		values := []int{9, 4, 2, 6, 8, 0, 3, 1, 7, 5}
		heapsort(values)
		ensureIntSlicesMatch(t, values, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	})

	t.Run("randomized", func(t *testing.T) {
		var iterations, max int
		if testing.Short() {
			iterations = 100
			max = 100
		} else {
			iterations = 1000
			max = 1000
		}

		for j := 0; j < iterations; j++ {
			want := make([]int, max)
			for i := 0; i < max; i++ {
				want[i] = i
			}

			values := rand.Perm(max) // generate a randomized list
			heapsort(values)

			ensureIntSlicesMatch(t, values, want)
		}
	})
}

func BenchmarkHeap(b *testing.B) {
	b.Run("100", func(b *testing.B) { benchSort(b, heapsort, 100) })
	b.Run("1000", func(b *testing.B) { benchSort(b, heapsort, 1000) })
	b.Run("10000", func(b *testing.B) { benchSort(b, heapsort, 10000) })
}
