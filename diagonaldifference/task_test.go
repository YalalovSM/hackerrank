package diagonaldifference

import "testing"

func TestTask(t *testing.T) {
	var arr [][]int32 = [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}

	if got := DiagonalDifference(arr); got != 15 {
		t.Errorf("diff = %d; want 15", got)
	}
}
