package main

import (
	"ASCII/Functions"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please enter correct usage: go run . \"Insert Text\"")
		return
	}

	// Define the path to the file
	filePath := "ART/standard.txt"
	// Open the file and handle any errors
	content, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Failed to read file:", err)
		return
	}
	defer content.Close()
	// Check if the input is valid
	checkInput := os.Args[1]
	if !Functions.IsValidInput(checkInput) {
		fmt.Println("Please enter valid characters")
		return
	}
	if checkInput == "\\n" {
		fmt.Println()
		return
	}
	// Split the input into lines

	userInput := strings.Split(checkInput, "\\n")
	if Functions.IsEmpty(userInput) {
		userInput = userInput[:len(userInput)-1]
	}
	Functions.PrintART(content, userInput)
}
