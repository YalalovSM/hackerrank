package comparetriplets

import "testing"

func TestTask(t *testing.T) {
	var a []int32 = []int32{5, 6, 7}
	var b []int32 = []int32{3, 6, 10}

	got := CompareTriplets(a, b)
	if got[0] != 1 {
		t.Errorf("a = %d; want 1", got[0])
	}

	if got[1] != 1 {
		t.Errorf("b = %d; want 1", got[1])
	}
}
