package animal

import "fmt"

type Animal interface {
	Sound()
}

type Dog struct {
	Name string
}

func (d Dog) Sound() {
	fmt.Println("Woof!")
}

type Cat struct {
	Name string
}

func (c Cat) Sound() {
	fmt.Println("Meow!")
}

func MakeSound(a Animal) {
	a.Sound()
}