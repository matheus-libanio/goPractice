package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Starting... ")

	file, err := os.Create("newCustomers.txt")
	if err != nil {
		defer panic(errors.New("The indicated file could not be created"))
	}
	defer file.Close()

	// Create a scanner to read input from the terminal
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter customer details. Type 'exit' to finish.")

	for {
		// Read the customer details from the terminal
		fmt.Print("Enter customer info: ")
		customer := make(map[string]string)
		// Input Name
		fmt.Print("Enter Name: ")
		scanner.Scan()
		customer["Name"] = scanner.Text()

		// Input ID
		fmt.Print("Enter ID: ")
		scanner.Scan()
		customer["ID"] = scanner.Text()

		// Input Phone Number
		fmt.Print("Enter Phone Number: ")
		scanner.Scan()
		customer["Phone Number"] = scanner.Text()

		// Input Address
		fmt.Print("Enter Address: ")
		scanner.Scan()
		customer["Address"] = scanner.Text()
		// Write customer details to the file
		formattedOutput := fmt.Sprintf("Name: %s, ID: %s, Phone Number: %s, Address: %s\n",
			customer["Name"], customer["ID"], customer["Phone Number"], customer["Address"])

		_, err := file.WriteString(formattedOutput)
		if err != nil {
			panic(err) // Panic if we can't write to the file
		}

		fmt.Println("Customer details saved. Enter another customer or type 'exit' to finish.")

		// Check if the user wants to exit
		fmt.Print("Do you want to add another customer? (y/n): ")
		scanner.Scan()
		if strings.ToLower(scanner.Text()) == "n" {
			break // Exit the loop if the user types 'n'
		}
	}

	fmt.Println("Customer details have been saved successfully.")

	fileReader, err := os.ReadFile("newCustomers.txt")
	if err != nil {
		defer panic(errors.New("The indicated file was empty"))
	}
	fmt.Println("Contents of file:\n", string(fileReader))
	fmt.Println("execução concluída")
}
