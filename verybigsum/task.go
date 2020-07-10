package verybigsum

// AVeryBigSum ...
// Link to the task: https://www.hackerrank.com/challenges/a-very-big-sum/problem
func AVeryBigSum(ar []int64) int64 {
	var sum int64 = 0

	for _, v := range ar {
		sum = sum + v
	}

	return sum
}
