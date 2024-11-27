package main

import (
	"fmt"

	tiposDeDados "github.com/matheus-libanio/goPractice/exec3/tipos_de_dados"
)

func main() {

	fmt.Println("exercicios 3 e 4")
	numero, err := tiposDeDados.RespostaDosExercicios()

	fmt.Println(numero, err)
}
