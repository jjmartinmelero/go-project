package main

func Sum(a, b int) int {
	return a + b
}

func GetBiggestNumber(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}
