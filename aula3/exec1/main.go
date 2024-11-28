package main

import (
	"fmt"

	calcimposto "github.com/matheus-libanio/goPractice/exec1/calc_Imposto"
)

func main() {

	salario := 157

	fmt.Printf("Salario de %vK\n", salario)
	imposto := calcimposto.CalcImposto(salario)

	fmt.Println("Imposto a ser deduzido do salario: ", imposto, "%")

}
