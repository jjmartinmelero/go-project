package main

import "fmt"

func main() {
	functionWithNparameters("Hello", 10, true)

	fmt.Println(factorial(3))

	// fn anonimas

	// 1:
	// func() {
	// 	fmt.Println("fn anonima")
	// }()

	// 2:
	greeting := func(name string) {
		fmt.Printf("fn anonima %s \n", name)
	}

	greeting("Juan Jesus")
	greetings("Juan Jesus", greeting)

	// fn as parameter
	result1 := aplicate(double, 2)
	result2 := aplicate(triple, 2)

	fmt.Printf("result1: %d, result2: %d\n", result1, result2)

	// orden superior
	fmt.Println("result", doubleV2(addOne, 3))

	// Closures
	nextInt := incrementer()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}

// fn Closures
func incrementer() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// fn orden superior
func doubleV2(f func(int) int, x int) int {
	return f(x * 2)
}

func addOne(x int) int {
	return x + 1
}

// fn with fn param ----------
func greetings(name string, f func(string)) {
	f(name)
}

func double(n int) int {
	return n * 2
}

func triple(n int) int {
	return n * 3
}

func aplicate(f func(int) int, n int) int {
	return f(n)
}

// fn with fn param ----------

// fn variadicas
func functionWithNparameters(data ...interface{}) {
	for _, dataType := range data {
		fmt.Printf("%T - %v\n", dataType, dataType)
	}
}

// fn recursiva
func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}
