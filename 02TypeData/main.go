package main

import "fmt"

// Package-level variable declaration
const Pi = 3.14
const (
	x = 100
	y = 0b1010
	z = 0o12
	w = 0xFF
)

const (
	Domingo = iota + 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main() {

	fmt.Println("\n=== Demostración de Variables y Constantes ===")
	// ************ SEVERAL OPTIONS TO DECLARE VARIABLES ************

	//	var firstName, lastName string
	//	var age int

	// var (
	// 	firstName = "Juan Jesus"
	// 	lastName = "Melero"
	// 	age = 31
	// )

	// var (
	// 	firstName, lastName string
	// 	age                 int
	// )

	// firstName = "Juan Jesus"
	// lastName = "Melero"
	// age = 31

	// var firstName, lastName, age = "Juan Jesus", "Melero", 31

	// Using the short declaration operator (:=)
	// only inside functions (not in package scope)
	firstName, lastName, age := "Juan Jesus", "Melero", 31

	fmt.Println("Hello, my name is", firstName, lastName, "and I am", age, "years old.")
	fmt.Println("Pi constant value is:", Pi)
	fmt.Println("Other constants:", x, y, z, w)
	fmt.Println("Days of the week:", Domingo, Lunes, Martes, Miercoles, Jueves, Viernes, Sabado)

	ShowBasicTypes()
	ShowTypeConversions()
	ShowFmtFeatures()
}
