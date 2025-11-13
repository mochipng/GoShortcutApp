package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(menu)
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		// trim whitespace
		input = strings.TrimSpace(input)

		// call function based on user input
		switch input {
		case "1":
			addShortcut()
		case "2":
			deleteShortcut()
		case "3":
			viewAllShortcuts()
		case "4":
			fmt.Println("Exiting program...")
			return
		default:
			fmt.Println("Invalid option, must be 1-4.")
			continue
		}
	}
}
