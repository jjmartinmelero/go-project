package main

import "testing"

func TestSum(test *testing.T) {
	// result := Sum(2, 3)
	// expect := 5

	// if result != expect {
	// 	test.Errorf("Error - expect value %d, received: %d", expect, result)
	// }

	cases := []struct {
		a, b, e int
	}{
		{3, 2, 5},
		{0, 0, 0},
		{-1, 1, 0},
	}

	for _, caseStruct := range cases {
		result := Sum(caseStruct.a, caseStruct.b)
		if result != caseStruct.e {
			test.Errorf("Error - expect value %d, received: %d", caseStruct.e, result)
		}
	}

}

func TestBiggestNumber(test *testing.T) {

	cases := []struct {
		a, b, e int
	}{
		{3, 2, 3},
		{0, 0, 0},
		{2, 7, 7},
	}

	for _, caseStruct := range cases {
		result := GetBiggestNumber(caseStruct.a, caseStruct.b)
		if result != caseStruct.e {
			test.Errorf("Error - expect value %d, received: %d", caseStruct.e, result)
		}
	}

}

func TestFibonacci(test *testing.T) {
	cases := []struct {
		n, e int
	}{
		{0, 0},
		{1, 1},
		{7, 13},
		{40, 102334155},
	}

	for _, caseStruct := range cases {
		result := Fibonacci(caseStruct.n)

		if result != caseStruct.e {
			test.Errorf("Error - expect value %d, received: %d", caseStruct.e, result)
		}
	}

}
