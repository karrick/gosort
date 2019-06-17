# gosort

sorting algorithms expressed in Go

## Observations

At 10 integers, insertion sort (2) is most efficient.

At 100 integers, insertion sort (3) is most efficient.

Somewhere between 1k and 10k integers, depending on the architecture
and OS, heap sort is more efficient than insertion sort (3).
