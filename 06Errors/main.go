package main

import (
	"fmt"
	"log"
	"os"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Error - Division by zero")
	}

	return a / b, nil
}

func dividePanic(a, b int) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	validateZero(b)

	fmt.Println(a / b)
}

func validateZero(b int) {
	if b == 0 {
		panic("Error - Division by zero")
	}
}

func main() {
	res, error := divide(10, 0)

	if error != nil {
		fmt.Println("Error: ", error)
		// return
	}

	fmt.Println("Result: ", res)

	// DEFERRED PANIC

	defer fmt.Println(3)
	defer fmt.Println(2)
	defer fmt.Println(1)

	file, err := os.Create("hola.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	defer file.Close()

	_, err = file.Write([]byte("Hola, Juan Jesus"))

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// Panic and Recover
	dividePanic(100, 10)
	dividePanic(200, 25)
	dividePanic(200, 0)
	dividePanic(100, 4)

	// Register error with log
	// log.Panic("This is a log panic message")
	// log.SetPrefix("LOG - main FN: ")
	// log.Println("This is an log message")
	// log.Println("This is another log message")
	// log.Fatal("This is a log fatal message")

	// Register log in a file
	logFile, errLogFile := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if errLogFile != nil {
		log.Fatal("Error opening log file", errLogFile)
	}

	defer logFile.Close()

	log.SetOutput(logFile)
	log.Println("This is a log message")
	log.Println("This is another log message")
	log.Fatal("This is a log fatal message" + " - Check the info.log file")

}
