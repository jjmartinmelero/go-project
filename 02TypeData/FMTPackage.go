package main

import (
	"fmt"
)

func ShowFmtFeatures() {
	fmt.Println("\n=== Funciones del paquete FMT ===")

	fmt.Print("Message without newline")

	name := "Juan"
	age := 32

	fmt.Printf("\n Hello my name is %s and I am %d years old.\n", name, age)

	greeting := fmt.Sprintf("\n Hello my name is %s and I am %d years old.", name, age)

	fmt.Printf("\n Hello my name is %v and I am %v years old.", name, age)
	
	fmt.Println(greeting)

	var newName, newName2 string
	var newAge int

	fmt.Print("\n Enter your name: ")
	fmt.Scanln(&newName, &newName2)

	fmt.Print("\n Enter your age: ")
	fmt.Scanln(&newAge)

	fmt.Println("Hello my name is", newName, newName2, "and I am", newAge, "years old.")

	// Checking types with %T
	fmt.Printf("Type of variable name is: %T\n", name)
	fmt.Printf("Type of variable age is: %T\n", age)

}
