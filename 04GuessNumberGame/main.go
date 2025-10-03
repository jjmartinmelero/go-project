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
		fmt.Println("Maniana !")
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

}
