package main

import (
	"fmt"
	"math/rand"
)

func startGame() {
	fmt.Println("Game !!")

	fmt.Println(rand.Intn(100))
	play()
}


func play(){
	randomNum := rand.Intn(100)
	var userNumber int
	var healths int
	const maxHealths = 10

	for healths < maxHealths {
		fmt.Printf("Ingresa un numero (Intentos restantes %d)", (maxHealths - healths))
		fmt.Scanln(&userNumber)

		if userNumber == randomNum {
			fmt.Println("Congrats ! ")
			tryAgain()
			return
		}
		
		if userNumber < randomNum {
			fmt.Println("your number is more less")
		} else {
			fmt.Println("your number is bigger")
		}

		healths++
	}


	fmt.Printf("\nGame over ! the guess number is %d \n", randomNum)
	tryAgain()
}

func tryAgain(){
	var userSelection string
	fmt.Println("Are you want play again ? s/n ")
	fmt.Scanln(&userSelection)

	switch userSelection {
	case "s": 
		play()
	case "n" :
		fmt.Println("Goodbye !")
	default:
		tryAgain() 

	}

}