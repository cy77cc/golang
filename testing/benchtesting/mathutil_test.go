package benchtesting

import "testing"

func BenchmarkAdd(b *testing.B) {
	b.ReportAllocs()
	/*
	测试用例				测试次数			平均每一次用时		在堆里开辟的内存		创建了几个对象
	BenchmarkAdd-8     1000000000	  0.3120 ns/op	    0 B/op	        	0 allocs/op
	*/
	for i :=0; i < b.N; i++ {
		Add(10, 20)
	}
}

func BenchmarkGetSum(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		GetSum(1000)
	}
}

func BenchmarkGetSumRecursively(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		GetSumRecursively(1000)
	}
}