package gosort

func quick(a []int) {
	quicka(a, 0, len(a)-1)
}

func quicka(a []int, lo, hi int) {
	if lo < hi {
		p := quickPivot(a, lo, hi)
		quicka(a, lo, p)
		quicka(a, p+1, hi)
	}
}

func quickPivot(a []int, lo, hi int) int {
	pivot := a[(lo+hi)>>1]
	i := lo - 1
	j := hi + 1

sortPivot:
findI:
	i++
	if a[i] < pivot {
		goto findI
	}
findJ:
	j--
	if a[j] > pivot {
		goto findJ
	}
	if i >= j {
		return j // done
	}
	t := a[j]
	a[j] = a[i]
	a[i] = t
	goto sortPivot
}
