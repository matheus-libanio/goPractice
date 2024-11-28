package main

import (
	"fmt"

	calculosalario "github.com/matheus-libanio/goPractice/exec3/calculo_salario"
)

func main() {
	var minutos int
	var categoria string

	fmt.Println("Entre a categoria que se enquadra: ")

	_, err := fmt.Scan(&categoria)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Entre o tempo trabalhado em minutos: ")

	_, err = fmt.Scan(&minutos)
	if err != nil {
		fmt.Println(err)
	}

	salario, err := calculosalario.CalculoSalario(minutos, categoria)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Salário a ser recebido é: US$", salario)

}
