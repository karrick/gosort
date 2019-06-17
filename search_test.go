package gosort

import "testing"

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
