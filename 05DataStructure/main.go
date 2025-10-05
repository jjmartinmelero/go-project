package main

import "fmt"

// Structs
type Person struct {
	name  string
	age   int
	email string
}

func main() {
	// var matriz = [5]int{10, 20, 30, 40, 50}
	var matriz = [...]int{10, 20, 30, 40, 50}
	matriz[0] = 100
	matriz[2] = 200

	fmt.Println(matriz)

	for i := 0; i < len(matriz); i++ {
		fmt.Println(matriz[i])
	}

	for index, value := range matriz {
		fmt.Printf("Indice: %d - Valor: %d \n", index, value)
	}

	// bidimensional
	var matriz2 = [2][3]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println(matriz2)

	//slice
	daysWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	dayFromWeek := daysWeek[:5]
	fmt.Println(dayFromWeek)

	dayFromWeek = append(dayFromWeek, "Funday")
	fmt.Println(dayFromWeek)

	//remove
	dayFromWeek = append(dayFromWeek[:2], dayFromWeek[3:]...)

	fmt.Println(dayFromWeek)
	fmt.Println(len(dayFromWeek))
	fmt.Println(cap(dayFromWeek))

	names := make([]string, 5, 10)
	names[0] = "John"
	fmt.Println(names)

	slice2 := []int{1, 2, 3, 4, 5}
	slice3 := make([]int, 2)

	copy(slice2, slice3)
	fmt.Println(slice2)
	fmt.Println(slice3)

	// maps
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	fmt.Println(colors)
	fmt.Println(colors["red"])

	colors["white"] = "#ffffff"

	fmt.Println(colors)
	// redValue := colors["red"]

	redValue, exist := colors["red"]
	fmt.Println(redValue, exist)

	blackValue, exist := colors["black"]

	if exist {
		fmt.Println("The color black exists")
	} else {
		fmt.Println("The color black does not exists")
	}

	fmt.Println(blackValue, exist)

	// remove an element
	delete(colors, "green")
	fmt.Println(colors)

	for key, value := range colors {
		fmt.Println(key, value)
	}

	// Structs

	var person Person
	person2 := Person{name: "Bob", age: 30, email: "test@google.com"}

	person.name = "Alice"
	person.age = 20
	person.email = "test@google.com"

	fmt.Println(person)
	fmt.Println(person2)

	// pointers

	var x int = 10
	fmt.Println(x)

	editValue(&x)

	fmt.Println(x)

	// var p *int = &x
	// fmt.Println(p)

	person.sayHello()

}

func editValue(x *int) {
	*x = 20
}

func (p *Person) sayHello() {
	fmt.Println("Hello, my name is ", p.name)
}
