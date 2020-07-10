package verybigsum

import "testing"

func TestTask(t *testing.T) {
	var arr []int64 = []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}

	if got := AVeryBigSum(arr); got != 5000000015 {
		t.Errorf("b = %d; want 5000000015", got)
	}

}
