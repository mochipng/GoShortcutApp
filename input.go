package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getInput(prompt string) string {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}

	return strings.TrimSpace(input)
}

func validateInput(name string) bool {
	return len(name) >= 2 && len(name) <= 20
}

func getValidateInput(prompt string, validate func(string) bool) string {
	for {
		input := getInput(prompt)
		if validate(input) {
			return input
		}

		fmt.Println("Invalid input. Please try again.")
	}
}
