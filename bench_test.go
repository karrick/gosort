package gosort

import (
	"math/rand"
	"testing"
)

func benchSort(b *testing.B, sorter func([]int), count int) {
	b.Helper()

	randomValues := rand.Perm(count)

	sortedValues := make([]int, count)
	for i := 0; i < count; i++ {
		sortedValues[i] = i
	}

	b.ResetTimer()

	for j := 0; j < b.N; j++ {
		values := make([]int, len(randomValues))
		copy(values, randomValues)
		sorter(values)
	}
}

func Benchmark10(b *testing.B) {
	const count = 10
	b.Run("heap", func(b *testing.B) { benchSort(b, heapsort, count) })
	b.Run("insertion1", func(b *testing.B) { benchSort(b, insertionsort1, count) })
	b.Run("insertion2", func(b *testing.B) { benchSort(b, insertionsort2, count) })
	b.Run("insertion3", func(b *testing.B) { benchSort(b, insertionsort3, count) })
}

func Benchmark50(b *testing.B) {
	const count = 50
	b.Run("heap", func(b *testing.B) { benchSort(b, heapsort, count) })
	b.Run("insertion1", func(b *testing.B) { benchSort(b, insertionsort1, count) })
	b.Run("insertion2", func(b *testing.B) { benchSort(b, insertionsort2, count) })
	b.Run("insertion3", func(b *testing.B) { benchSort(b, insertionsort3, count) })
}

func Benchmark100(b *testing.B) {
	const count = 100
	b.Run("heap", func(b *testing.B) { benchSort(b, heapsort, count) })
	b.Run("insertion1", func(b *testing.B) { benchSort(b, insertionsort1, count) })
	b.Run("insertion2", func(b *testing.B) { benchSort(b, insertionsort2, count) })
	b.Run("insertion3", func(b *testing.B) { benchSort(b, insertionsort3, count) })
}

func Benchmark1K(b *testing.B) {
	const count = 1000
	b.Run("heap", func(b *testing.B) { benchSort(b, heapsort, count) })
	b.Run("insertion1", func(b *testing.B) { benchSort(b, insertionsort1, count) })
	b.Run("insertion2", func(b *testing.B) { benchSort(b, insertionsort2, count) })
	b.Run("insertion3", func(b *testing.B) { benchSort(b, insertionsort3, count) })
}

func Benchmark5K(b *testing.B) {
	const count = 5000
	b.Run("heap", func(b *testing.B) { benchSort(b, heapsort, count) })
	b.Run("insertion1", func(b *testing.B) { benchSort(b, insertionsort1, count) })
	b.Run("insertion2", func(b *testing.B) { benchSort(b, insertionsort2, count) })
	b.Run("insertion3", func(b *testing.B) { benchSort(b, insertionsort3, count) })
}

func Benchmark10K(b *testing.B) {
	const count = 10000
	b.Run("heap", func(b *testing.B) { benchSort(b, heapsort, count) })
	b.Run("insertion1", func(b *testing.B) { benchSort(b, insertionsort1, count) })
	b.Run("insertion2", func(b *testing.B) { benchSort(b, insertionsort2, count) })
	b.Run("insertion3", func(b *testing.B) { benchSort(b, insertionsort3, count) })
}

func Benchmark100K(b *testing.B) {
	const count = 100000
	b.Run("heap", func(b *testing.B) { benchSort(b, heapsort, count) })
	// b.Run("insertion1", func(b *testing.B) { benchSort(b, insertionsort1, count) })
	// b.Run("insertion2", func(b *testing.B) { benchSort(b, insertionsort2, count) })
	// b.Run("insertion3", func(b *testing.B) { benchSort(b, insertionsort3, count) })
}

func Benchmark1M(b *testing.B) {
	const count = 1000000
	b.Run("heap", func(b *testing.B) { benchSort(b, heapsort, count) })
	// b.Run("insertion1", func(b *testing.B) { benchSort(b, insertionsort1, count) })
	// b.Run("insertion2", func(b *testing.B) { benchSort(b, insertionsort2, count) })
	// b.Run("insertion3", func(b *testing.B) { benchSort(b, insertionsort3, count) })
}
