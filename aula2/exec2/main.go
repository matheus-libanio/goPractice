package main

import (
	"errors"
	"fmt"
)

func main() {

	var idade, tempoContribuicao, salario int

	fmt.Println("Enter your age: ")

	_, err := fmt.Scan(&idade)
	if err != nil {
		errors.New("Error entering age")
	}

	fmt.Println("Enter your time of contributtion: ")

	_, err = fmt.Scan(&tempoContribuicao)
	if err != nil {
		errors.New("Error entering time of contribution")
	}

	fmt.Println("Enter your salary: ")

	_, err = fmt.Scan(&salario)
	if err != nil {
		errors.New("Error entering salary")
	}

	if idade < 22 || tempoContribuicao < 1 || salario < 100000 {
		fmt.Println("you are not ellegible for credits")
		switch {
		case idade < 22:
			fmt.Println("Too young")
		case tempoContribuicao < 1:
			fmt.Println("little contribution")
		default:
			fmt.Println("Too poor")
		}
	} else {
		fmt.Println("you are ellegible for credits, congrats!")
	}

}
