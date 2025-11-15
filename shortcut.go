package main

import (
	"fmt"
)

// storing shortcuts
var shortcuts = map[string]string{}

func findShortcut(name string) (string, bool) {
	for code, keys := range shortcuts {
		if name == keys {
			return code, true
		}
	}
	return "", false
}

func addShortcut(name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("Cannot add: cannot be empty")
	}

	code := generateCode()
	shortcuts[code] = name

	fmt.Printf("New shortcut created: %s", name)
	return code, nil
}

func deleteShortcut(name string, validate func(string) bool) error {
	if !validate(name) {
		return fmt.Errorf("Error: shortcut name not valid")
	}

	code, found := findShortcut(name)
	if !found {
		return fmt.Errorf("Error: shortcut not found")
	}

	delete(shortcuts, code)
	return nil

}

func viewAllShortcuts() string {}

func generateCode() string {
	number := len(shortcuts) + 1
	return fmt.Sprintf("%03d", number)
}
