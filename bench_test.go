package gosort

import (
	"flag"
	"math/rand"
	"os"
	"testing"
)

var randomValues []int
var sortedValues []int

func TestMain(m *testing.M) {
	flag.Parse()

	var max int
	if testing.Short() {
		max = 1000
	} else {
		max = 10000
	}

	randomValues = rand.Perm(max)

	sortedValues = make([]int, max)
	for i := 0; i < max; i++ {
		sortedValues[i] = i
	}

	os.Exit(m.Run())
}

func benchSort(b *testing.B, sorter func([]int), count int) {
	// 	b.Helper()

	// 	for j := 0; j < b.N; j++ {
	// 		values := make([]int, len(randomValues))
	// 		copy(values, randomValues)
	// 		sorter(values)
	// 		ensureIntSlicesMatch(b, values, sortedValues)
	// 	}
	// }

	// func benchSort2(b *testing.B, sorter func([]int), count int) {
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
		ensureIntSlicesMatch(b, values, sortedValues)
	}
}
