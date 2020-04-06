package main

func expand(val int, vals []int) []int {
	nvals := make([]int, 0)
	for _, v := range vals {
		nvals = append(nvals, v)
	}

	if val == 4 {
		nvals = append(nvals, 5)
	}

	if val == 3 {
		nvals = append(nvals, []int{4, 5}...)
	}

	return nvals
}

func solve151(depth int, vals []int) (int, int) {
	if len(vals) == 0 {
		return 0, 0
	}

	total := 0
	singles := 0

	if depth > 0 && depth < 3 {
		if len(vals) == 1 {
			singles++
		}

		total++
	}

	for i, v := range vals {
		tvals := make([]int, 0)
		for j := range vals {
			if j != i {
				tvals = append(tvals, vals[j])
			}
		}

		t, s := solve151(depth+1, expand(v, tvals))
		total += t
		singles += s
	}

	return total, singles
}
