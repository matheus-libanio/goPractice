package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Starting... ")
	_, err := os.Open("customers.txt")
	if err != nil {
		panic(errors.New("The indicated file was not found or is damaged"))
	}
	fmt.Println("End")
 }
 