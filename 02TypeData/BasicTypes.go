package main

import "fmt"

func ShowBasicTypes() {
	fmt.Println("\n=== Tipos de Datos Básicos ===")
	fullName := "Juan Jesus Melero \t(alias \"jj\")\n"
	fmt.Println(fullName)

	var a byte = 'a'
	fmt.Println(a)

	s := "Hello, Go!"
	fmt.Println(s[0])

	var r rune = 'ñ'
	fmt.Println(r)

	// values by default (zero values)
	var (
		defaultInt    int
		defaultUint   uint
		defaultFloat  float32
		defaultBool   bool
		defaultString string
	)

	fmt.Printf("Default int: %v, uint: %v, float32: %v, bool: %v, string: '%v'\n", defaultInt, defaultUint, defaultFloat, defaultBool, defaultString)

}
