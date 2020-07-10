package diagonaldifference

// DiagonalDifference ...
// Link to the task: https://www.hackerrank.com/challenges/diagonal-difference/problem
func DiagonalDifference(arr [][]int32) int32 {
	len := len(arr[0])

	var leftToRight int32 = 0
	for i := 0; i < len; i++ {
		leftToRight = leftToRight + arr[i][i]
	}

	var rightToLeft int32 = 0
	for i := 0; i < len; i++ {
		rightToLeft = rightToLeft + arr[i][len-i-1]
	}

	return abs(leftToRight - rightToLeft)
}

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}
