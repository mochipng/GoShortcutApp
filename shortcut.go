package main

import (
	"errors"
	"fmt"
)

// storing shortcuts
var shortcuts = map[string]string{}
var shortcutCounter = 0

func findShortcut(name string) (string, bool) {
	for code, shortcutName := range shortcuts {
		if name == shortcutName {
			return code, true
		}
	}
	return "", false
}

func hasShortcuts() bool {
	return len(shortcuts) > 0
}

func generateCode() string {
	shortcutCounter++
	return fmt.Sprintf("%03d", shortcutCounter)
}

func addShortcut(name string) (string, error) {
	code := generateCode()
	shortcuts[code] = name
	return code, nil
}

func deleteShortcut(name string) error {
	code, found := findShortcut(name)
	if !found {
		return errors.New("shortcut not found")
	}

	delete(shortcuts, code)
	return nil

}

func viewAllShortcuts() {
	fmt.Println("All shortcuts:")
	for code, name := range shortcuts {
		fmt.Printf("%s -> %s\n", code, name)
	}
}
