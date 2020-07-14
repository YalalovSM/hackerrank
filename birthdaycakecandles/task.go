package birthdaycakecandles

// BirthdayCakeCandles ...
func BirthdayCakeCandles(ar []int32) int32 {
	var max int32 = 0
	count := 0

	for i := 0; i < len(ar); i++ {
		if ar[i] > max {
			max = ar[i]
		}
	}

	for i := 0; i < len(ar); i++ {
		if ar[i] == max {
			count++
		}
	}

	return int32(count)
}
