package benchtesting

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

func Div(a, b int) int {
	return a / b
}

func GetSum(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func GetSumRecursively(n int) int {
	if n == 1 {
		return 1
	}
	return n + GetSumRecursively(n-1)
}
