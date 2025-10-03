package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	if t := time.Now(); t.Hour() < 12 {
		fmt.Println("Maniana !")
	} else if t.Hour() < 17 {
		fmt.Println("tarde !")
	} else {
		fmt.Println("Noche !")
	}

	switch t := time.Now(); {
	case t.Hour() < 12:
		fmt.Println(t.Hour(), "Maniana !")
	case t.Hour() < 17:
		fmt.Println("Tarde !")
	default:
		fmt.Println("Noche !")
	}

	os := runtime.GOOS

	switch os {
	case "windows":
		fmt.Println("Go run -> Windows")
	case "linux":
		fmt.Println("Go run -> Linux")
	case "darwin":
		fmt.Println("Go run -> Mac OS")
	default:
		fmt.Println("Go run -> Otro OS")
	}

	// var i int


	for  i := 1; i < 10; i++ {
		
		if i == 5 {
			continue
		}

		fmt.Println(i)
	}

	greeting := hello("Juan Jesus")
	fmt.Println(greeting)

	sum, mul := calc(2, 4)

	fmt.Printf("sum : %d and mul: %d \n", sum, mul)

	startGame()
}

func hello(name string) string {
	fmt.Printf("Hello %s from func! \n", name)
	return "Hello, " + name
}

func calc(a, b int) (sum, mul int){

	// sum := a + b
	sum = a + b
	// mul := a * b
	mul = a * b

	// return sum, mul
	return
}