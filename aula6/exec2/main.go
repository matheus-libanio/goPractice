package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Starting... ")
	file, err := os.Open("customers.txt")
	if err != nil {
		defer panic(errors.New("The indicated file was not found or is damaged"))
	}
	defer file.Close()

	fileReader, err := os.ReadFile("customers.txt")
	if err != nil {
		defer panic(errors.New("The indicated file was empty"))
	}
	fmt.Println("Contents of file:\n", string(fileReader))
	fmt.Println("execução concluída")
}
