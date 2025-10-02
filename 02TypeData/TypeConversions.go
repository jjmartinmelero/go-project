package main

import (
	"fmt"
	"strconv"
)

func ShowTypeConversions() {
	fmt.Println("\n=== Conversiones entre Tipos Numéricos ===")
	var integer16 int16 = 50
	var integer32 int32 = 100

	fmt.Println(int32(integer16) + integer32)

	s := "100"
	i, _ := strconv.Atoi(s)

	fmt.Println(i + i)

	n := 42
	s = strconv.Itoa(n)
	fmt.Println(s + s)
}
