package main

import (
	"errors"
	"fmt"
)

// storing shortcuts
var shortcuts = map[string]string{}

func findShortcut(name string) (string, bool) {
	for code, value := range shortcuts {
		if name == value {
			return code, true
		}
	}
	return "", false
}

func addShortcut(name string) (string, error) {
	if name == "" {
		return "", errors.New("Error: cannot be empty")
	}

	code := generateCode()
	shortcuts[code] = name
	return code, nil
}

func deleteShortcut(name string, validate func(string) bool) error {
	if !validate(name) {
		return errors.New("Error: shortcut name not valid")
	}

	code, found := findShortcut(name)
	if !found {
		return errors.New("Error: shortcut not found")
	}

	delete(shortcuts, code)
	return nil

}

func viewAllShortcuts() {
	if len(shortcuts) == 0 {
		fmt.Println("No shortcuts added.")
		return
	}

	fmt.Println("All shortcuts:")
	for code, name := range shortcuts {
		fmt.Printf("%s -> %s\n", code, name)
	}
}

func generateCode() string {
	number := len(shortcuts) + 1
	return fmt.Sprintf("%03d", number)
}
