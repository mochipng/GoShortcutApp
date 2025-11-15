package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getInput(prompt string) string {
	for {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		} else {
			return strings.TrimSpace(input)
		}
	}
}

func validateInput(prompt string, validate func(string) bool) string {
	for {
		input := getInput(prompt)
		if validate(input) {
			return input
		}

		fmt.Println("Invalid input, please try again.")
	}
}
