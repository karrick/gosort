# gosort

sorting algorithms expressed in Go

## Observations

At 10 integers, insertion sort (2) is most efficient.

At 100 integers, insertion sort (3) is most efficient.

At about 5k integers, insertion sort (3) is slightly faster than heap
sort.

At about 7500, heap sort is more efficient than insertion sort (3).
