package comparetriplets

// CompareTriplets ...
func CompareTriplets(a []int32, b []int32) []int32 {
	var r []int32

	r = make([]int32, 2)

	for i, v := range a {
		if v > b[i] {
			r[0] = r[0] + 1
		} else if v < b[i] {
			r[1] = r[1] + 1
		}
	}

	return r
}
