package mysort

import (
	"testing"
)

func BenchmarkCheckSort(b *testing.B) {
	b.ReportAllocs()
	ints := []int{10, 2, 3, 7, 9, 6, 4, 1, 5}
	for i := 0; i < b.N; i++ {
		CheckSort(ints)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	b.ReportAllocs()
	ints := []int{10, 2, 3, 7, 9, 6, 4, 1, 5}
	for i := 0; i < b.N; i++ {
		BubbleSort(ints)
	}
}