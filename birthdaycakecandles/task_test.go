package birthdaycakecandles

import "testing"

func TestTask(t *testing.T) {
	var ar []int32 = []int32{3, 2, 1, 3}

	got := BirthdayCakeCandles(ar)

	if got != 2 {
		t.Errorf("result = %d; want 2", got)
	}
}
