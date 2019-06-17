package gosort

import (
	"flag"
	"math/rand"
	"os"
	"testing"
)

func TestInsertion1(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		var values []int
		insertionsort1(values)
		ensureIntSlicesMatch(t, values, nil)
	})

	t.Run("a", func(t *testing.T) {
		values := []int{13}
		insertionsort1(values)
		ensureIntSlicesMatch(t, values, []int{13})
	})

	t.Run("two", func(t *testing.T) {
		t.Run("ab", func(t *testing.T) {
			values := []int{13, 42}
			insertionsort1(values)
			ensureIntSlicesMatch(t, values, []int{13, 42})
		})
		t.Run("ba", func(t *testing.T) {
			values := []int{42, 13}
			insertionsort1(values)
			ensureIntSlicesMatch(t, values, []int{13, 42})
		})
	})

	t.Run("three", func(t *testing.T) {
		t.Run("abc", func(t *testing.T) {
			values := []int{13, 42, 97}
			insertionsort1(values)
			ensureIntSlicesMatch(t, values, []int{13, 42, 97})
		})
		t.Run("bca", func(t *testing.T) {
			values := []int{42, 97, 13}
			insertionsort1(values)
			ensureIntSlicesMatch(t, values, []int{13, 42, 97})
		})
		t.Run("cab", func(t *testing.T) {
			values := []int{97, 13, 42}
			insertionsort1(values)
			ensureIntSlicesMatch(t, values, []int{13, 42, 97})
		})
		t.Run("acb", func(t *testing.T) {
			values := []int{13, 97, 42}
			insertionsort1(values)
			ensureIntSlicesMatch(t, values, []int{13, 42, 97})
		})
		t.Run("cba", func(t *testing.T) {
			values := []int{97, 42, 13}
			insertionsort1(values)
			ensureIntSlicesMatch(t, values, []int{13, 42, 97})
		})
		t.Run("bac", func(t *testing.T) {
			values := []int{42, 13, 97}
			insertionsort1(values)
			ensureIntSlicesMatch(t, values, []int{13, 42, 97})
		})
		t.Run("acb", func(t *testing.T) {
			values := []int{13, 97, 42}
			insertionsort1(values)
			ensureIntSlicesMatch(t, values, []int{13, 42, 97})
		})
	})

	t.Run("ten", func(t *testing.T) {
		values := []int{9, 4, 2, 6, 8, 0, 3, 1, 7, 5}
		insertionsort1(values)
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
			insertionsort1(values)

			ensureIntSlicesMatch(t, values, want)
		}
	})
}

func TestInsertion2(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		var values []int
		insertionsort2(values)
		ensureIntSlicesMatch(t, values, nil)
	})

	t.Run("a", func(t *testing.T) {
		values := []int{13}
		insertionsort2(values)
		ensureIntSlicesMatch(t, values, []int{13})
	})

	t.Run("two", func(t *testing.T) {
		t.Run("ab", func(t *testing.T) {
			values := []int{13, 42}
			insertionsort2(values)
			ensureIntSlicesMatch(t, values, []int{13, 42})
		})
		t.Run("ba", func(t *testing.T) {
			values := []int{42, 13}
			insertionsort2(values)
			ensureIntSlicesMatch(t, values, []int{13, 42})
		})
	})

	t.Run("three", func(t *testing.T) {
		t.Run("abc", func(t *testing.T) {
			values := []int{13, 42, 97}
			insertionsort2(values)
			ensureIntSlicesMatch(t, values, []int{13, 42, 97})
		})
		t.Run("bca", func(t *testing.T) {
			values := []int{42, 97, 13}
			insertionsort2(values)
			ensureIntSlicesMatch(t, values, []int{13, 42, 97})
		})
		t.Run("cab", func(t *testing.T) {
			values := []int{97, 13, 42}
			insertionsort2(values)
			ensureIntSlicesMatch(t, values, []int{13, 42, 97})
		})
		t.Run("acb", func(t *testing.T) {
			values := []int{13, 97, 42}
			insertionsort2(values)
			ensureIntSlicesMatch(t, values, []int{13, 42, 97})
		})
		t.Run("cba", func(t *testing.T) {
			values := []int{97, 42, 13}
			insertionsort2(values)
			ensureIntSlicesMatch(t, values, []int{13, 42, 97})
		})
		t.Run("bac", func(t *testing.T) {
			values := []int{42, 13, 97}
			insertionsort2(values)
			ensureIntSlicesMatch(t, values, []int{13, 42, 97})
		})
		t.Run("acb", func(t *testing.T) {
			values := []int{13, 97, 42}
			insertionsort2(values)
			ensureIntSlicesMatch(t, values, []int{13, 42, 97})
		})
	})

	t.Run("ten", func(t *testing.T) {
		t.Run("unsorted", func(t *testing.T) {
			values := []int{9, 4, 2, 6, 8, 0, 3, 1, 7, 5}
			insertionsort2(values)
			ensureIntSlicesMatch(t, values, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
		})
		t.Run("sorted", func(t *testing.T) {
			values := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
			insertionsort2(values)
			ensureIntSlicesMatch(t, values, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
		})
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
			insertionsort2(values)

			ensureIntSlicesMatch(t, values, want)
		}
	})
}

func TestInsertion3(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		var values []int
		insertionsort3(values)
		ensureIntSlicesMatch(t, values, nil)
	})

	t.Run("a", func(t *testing.T) {
		values := []int{13}
		insertionsort3(values)
		ensureIntSlicesMatch(t, values, []int{13})
	})

	t.Run("two", func(t *testing.T) {
		t.Run("ab", func(t *testing.T) {
			values := []int{13, 42}
			insertionsort3(values)
			ensureIntSlicesMatch(t, values, []int{13, 42})
		})
		t.Run("ba", func(t *testing.T) {
			values := []int{42, 13}
			insertionsort3(values)
			ensureIntSlicesMatch(t, values, []int{13, 42})
		})
	})

	t.Run("three", func(t *testing.T) {
		t.Run("abc", func(t *testing.T) {
			values := []int{13, 42, 97}
			insertionsort3(values)
			ensureIntSlicesMatch(t, values, []int{13, 42, 97})
		})
		t.Run("bca", func(t *testing.T) {
			values := []int{42, 97, 13}
			insertionsort3(values)
			ensureIntSlicesMatch(t, values, []int{13, 42, 97})
		})
		t.Run("cab", func(t *testing.T) {
			values := []int{97, 13, 42}
			insertionsort3(values)
			ensureIntSlicesMatch(t, values, []int{13, 42, 97})
		})
		t.Run("acb", func(t *testing.T) {
			values := []int{13, 97, 42}
			insertionsort3(values)
			ensureIntSlicesMatch(t, values, []int{13, 42, 97})
		})
		t.Run("cba", func(t *testing.T) {
			values := []int{97, 42, 13}
			insertionsort3(values)
			ensureIntSlicesMatch(t, values, []int{13, 42, 97})
		})
		t.Run("bac", func(t *testing.T) {
			values := []int{42, 13, 97}
			insertionsort3(values)
			ensureIntSlicesMatch(t, values, []int{13, 42, 97})
		})
		t.Run("acb", func(t *testing.T) {
			values := []int{13, 97, 42}
			insertionsort3(values)
			ensureIntSlicesMatch(t, values, []int{13, 42, 97})
		})
	})

	t.Run("ten", func(t *testing.T) {
		t.Run("unsorted", func(t *testing.T) {
			values := []int{9, 4, 2, 6, 8, 0, 3, 1, 7, 5}
			insertionsort3(values)
			ensureIntSlicesMatch(t, values, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
		})
		t.Run("partially sorted", func(t *testing.T) {
			values := []int{0, 1, 2, 3, 5, 4, 6, 7, 8, 9}
			insertionsort3(values)
			ensureIntSlicesMatch(t, values, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
		})
		t.Run("already sorted", func(t *testing.T) {
			values := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
			insertionsort3(values)
			ensureIntSlicesMatch(t, values, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
		})
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
			insertionsort3(values)

			ensureIntSlicesMatch(t, values, want)
		}
	})
}

func TestSearchLeftMostGreaterThan(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		if got, want := searchLeftMostGreaterThan(0, []int{0, 1, 2, 3, 4, 5, 6}), 1; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := searchLeftMostGreaterThan(1, []int{0, 1, 2, 3, 4, 5, 6}), 2; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := searchLeftMostGreaterThan(2, []int{0, 1, 2, 3, 4, 5, 6}), 3; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := searchLeftMostGreaterThan(3, []int{0, 1, 2, 3, 4, 5, 6}), 4; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := searchLeftMostGreaterThan(4, []int{0, 1, 2, 3, 4, 5, 6}), 5; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := searchLeftMostGreaterThan(5, []int{0, 1, 2, 3, 4, 5, 6}), 6; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := searchLeftMostGreaterThan(6, []int{0, 1, 2, 3, 4, 5, 6}), 7; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
		if got, want := searchLeftMostGreaterThan(7, []int{0, 1, 2, 3, 4, 5, 6}), 7; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
	t.Run("with a run", func(t *testing.T) {
		if got, want := searchLeftMostGreaterThan(3, []int{0, 1, 2, 3, 3, 3, 6}), 6; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})
}

func TestInsertion4(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		var values []int
		insertionsort4(values)
		ensureIntSlicesMatch(t, values, nil)
	})

	t.Run("a", func(t *testing.T) {
		values := []int{13}
		insertionsort4(values)
		ensureIntSlicesMatch(t, values, []int{13})
	})

	t.Run("two", func(t *testing.T) {
		t.Run("ab", func(t *testing.T) {
			values := []int{13, 42}
			insertionsort4(values)
			ensureIntSlicesMatch(t, values, []int{13, 42})
		})
		t.Run("ba", func(t *testing.T) {
			values := []int{42, 13}
			insertionsort4(values)
			ensureIntSlicesMatch(t, values, []int{13, 42})
		})
	})

	t.Run("three", func(t *testing.T) {
		t.Run("abc", func(t *testing.T) {
			values := []int{13, 42, 97}
			insertionsort4(values)
			ensureIntSlicesMatch(t, values, []int{13, 42, 97})
		})
		t.Run("bca", func(t *testing.T) {
			values := []int{42, 97, 13}
			insertionsort4(values)
			ensureIntSlicesMatch(t, values, []int{13, 42, 97})
		})
		t.Run("cab", func(t *testing.T) {
			values := []int{97, 13, 42}
			insertionsort4(values)
			ensureIntSlicesMatch(t, values, []int{13, 42, 97})
		})
		t.Run("acb", func(t *testing.T) {
			values := []int{13, 97, 42}
			insertionsort4(values)
			ensureIntSlicesMatch(t, values, []int{13, 42, 97})
		})
		t.Run("cba", func(t *testing.T) {
			values := []int{97, 42, 13}
			insertionsort4(values)
			ensureIntSlicesMatch(t, values, []int{13, 42, 97})
		})
		t.Run("bac", func(t *testing.T) {
			values := []int{42, 13, 97}
			insertionsort4(values)
			ensureIntSlicesMatch(t, values, []int{13, 42, 97})
		})
		t.Run("acb", func(t *testing.T) {
			values := []int{13, 97, 42}
			insertionsort4(values)
			ensureIntSlicesMatch(t, values, []int{13, 42, 97})
		})
	})

	t.Run("ten", func(t *testing.T) {
		t.Run("unsorted", func(t *testing.T) {
			values := []int{9, 4, 2, 6, 8, 0, 3, 1, 7, 5}
			insertionsort4(values)
			ensureIntSlicesMatch(t, values, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
		})
		t.Run("partially sorted", func(t *testing.T) {
			values := []int{0, 1, 2, 3, 5, 4, 6, 7, 8, 9}
			insertionsort4(values)
			ensureIntSlicesMatch(t, values, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
		})
		t.Run("already sorted", func(t *testing.T) {
			values := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
			insertionsort4(values)
			ensureIntSlicesMatch(t, values, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
		})
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
			insertionsort4(values)

			ensureIntSlicesMatch(t, values, want)
		}
	})
}

var randomValues []int
var sortedValues []int

func TestMain(m *testing.M) {
	flag.Parse()

	var max int
	if testing.Short() {
		max = 100
	} else {
		max = 1000
	}

	randomValues = rand.Perm(max)

	sortedValues = make([]int, max)
	for i := 0; i < max; i++ {
		sortedValues[i] = i
	}

	os.Exit(m.Run())
}

func BenchmarkInsertion1(b *testing.B) {
	for j := 0; j < b.N; j++ {
		values := make([]int, len(randomValues))
		copy(values, randomValues)
		insertionsort1(values)
		ensureIntSlicesMatch(b, values, sortedValues)
	}
}

func BenchmarkInsertion2(b *testing.B) {
	for j := 0; j < b.N; j++ {
		values := make([]int, len(randomValues))
		copy(values, randomValues)
		insertionsort2(values)
		ensureIntSlicesMatch(b, values, sortedValues)
	}
}

func BenchmarkInsertion3(b *testing.B) {
	for j := 0; j < b.N; j++ {
		values := make([]int, len(randomValues))
		copy(values, randomValues)
		insertionsort3(values)
		ensureIntSlicesMatch(b, values, sortedValues)
	}
}

func BenchmarkInsertion4(b *testing.B) {
	for j := 0; j < b.N; j++ {
		values := make([]int, len(randomValues))
		copy(values, randomValues)
		insertionsort4(values)
		ensureIntSlicesMatch(b, values, sortedValues)
	}
}
