package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type integer int

type Numbers interface {
	~int | ~float64 | ~float32
}

func main() {
	PrintList("test", "test2")

	fmt.Println(sum(1, 2, 3, 4, 5.75))

	// restricciones de aproximacion ~
	var num1 integer = 100
	var num2 integer = 200

	fmt.Println(sum(num1, num2))

	// slices
	strings := []string{"a", "b", "c", "d"}
	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println(includes(strings, "a"))
	fmt.Println(includes(strings, "f"))
	fmt.Println(includes(numbers, 2))
	fmt.Println(includes(numbers, 9))

	// filter fn
	fmt.Println(filter(numbers, func(value int) bool { return value >= 3 }))
	fmt.Println(filter(strings, func(value string) bool { return value > "b" }))

	// Products
	product1 := Product[uint]{1, "Zapatos", 50}
	product2 := Product[string]{"id1234", "Zapatos", 50}

	fmt.Println(product1)
	fmt.Println(product2)
}

// Generic Structure
type Product[T uint | string] struct {
	Id    T
	Desc  string
	Price float32
}

func filter[T constraints.Ordered](list []T, callback func(T) bool) []T {
	result := make([]T, 0, len(list))

	for _, item := range list {
		if callback(item) {
			result = append(result, item)
		}
	}

	return result
}

// With comparable type
func includes[T comparable](list []T, value T) bool {

	for _, item := range list {

		if item == value {
			return true
		}
	}

	return false
}

func sum[T constraints.Integer | constraints.Float](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}

	return total
}

func PrintList(list ...any) {
	for _, value := range list {
		fmt.Println(value)
	}
}
