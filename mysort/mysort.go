package mysort

func CheckSort(ints []int) {
	for i := 0; i < len(ints)-1; i++ {
		for j := i + 1; j < len(ints); j++ {
			if ints[i] > ints[j] {
				ints[i], ints[j] = ints[j], ints[i]
			}
		}
	}
}

func BubbleSort(ints []int) {
	for i := 1; i < len(ints); i++ {
		for j := 0; j < len(ints)-i; j++ {
			if ints[j] > ints[j+1] {
				ints[j], ints[j+1] = ints[j+1], ints[j]
			}
		}
	}
}
