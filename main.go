package main

import (
	"fmt"
)

func main() {
	// create menu and buffered reader
	menu := `
	SHORTCUT CREATOR
	------------------------
	Choose an option (1-4):
	1. Add
	2. Delete
	3. View All
	4. Exit
	`

	for {
		fmt.Println(menu)
		input := getInput("> ")

		// call function based on user input
		switch input {
		case "1":
			name := getValidateInput("Enter shortcut name: ", validateInput)
			code, err := addShortcut(name)

			if err != nil {
				fmt.Print("Error:", err)
			} else {
				fmt.Printf("Shortcut created. Code: %s, Action: %s\n", code, name)
			}

		case "2":
			name := getValidateInput("Enter shortcut name: ", validateInput)
			err := deleteShortcut(name)

			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Shortcut deleted.")
			}

		case "3":
			if !hasShortcuts() {
				fmt.Println("No shortcuts added.")
			} else {
				viewAllShortcuts()
			}

		case "4":
			fmt.Println("Exiting program...")
			return
		default:
			fmt.Println("Invalid option, must be 1-4.")
		}
	}
}
