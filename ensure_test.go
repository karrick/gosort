package gosort

import "testing"

func ensureIntSlicesMatch(tb testing.TB, actual, expected []int) {
	tb.Helper()

	la := len(actual)
	lb := len(expected)

	// iterate to higher of both indexes
	l := la
	if l < lb {
		l = lb
	}

	for i := 0; i < l; i++ {
		switch {
		case i > la:
			tb.Errorf("%d: WANT: %v (missing)", i, expected[i])
		case i > lb:
			tb.Errorf("%d: GOT: %v (extra)", i, actual[i])
		default:
			if got, want := actual[i], expected[i]; got != want {
				tb.Errorf("%d: GOT: %v; WANT: %v", i, got, want)
			}
		}
	}
	if tb.Failed() {
		tb.Log(actual)
	}
}
