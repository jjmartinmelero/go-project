package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Contact Structure
type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func saveContactsToFile(contacts []Contact) error {
	file, err := os.Create("contacts.json")

	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(contacts)

	if err != nil {
		return err
	}

	return nil
}

func loadContactsFromFile(contacts *[]Contact) error {
	file, err := os.Open("contacts.json")

	if err != nil {
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(file)

	if err != nil {
		return err
	}

	return nil
}

func main() {

	var contacts []Contact

	err := loadContactsFromFile(&contacts)

	if err != nil {
		fmt.Println("Error loading contacts: ", err)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("1. Add Contact",
			"\n2. View Contacts",
			"\n3. Exit")

		var option int
		_, err := fmt.Scanln(&option)

		if err != nil {
			fmt.Println("Invalid input")
			return
		}

		switch option {

		case 1:
			var c Contact

			fmt.Println("Enter Name: ")
			c.Name, _ = reader.ReadString('\n')
			fmt.Println("Enter Email: ")
			c.Email, _ = reader.ReadString('\n')
			fmt.Println("Enter Phone: ")
			c.Phone, _ = reader.ReadString('\n')

			contacts = append(contacts, c)

			if err := saveContactsToFile(contacts); err != nil {
				fmt.Println("Error saving contacts: ", err)
			}
		case 2:
			for i, c := range contacts {
				fmt.Printf("%d. %s - %s - %s\n", i+1, c.Name, c.Email, c.Phone)
			}
		case 3:
			fmt.Println("Exitings...")
			return
		default:
			fmt.Println("Invalid option")
		}

	}

}
