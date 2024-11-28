package main

import (
	"fmt"

	"github.com/matheus-libanio/goPractice/exec4/estatistica"
)

func main() {

	var categoria int

	fmt.Println("Entre a operação que deseja:\n1- Minumum\n2- Average\n3- Maximum")

	_, err := fmt.Scan(&categoria)
	if err != nil {
		fmt.Println(err)
	}

	funcSelecionada, err := estatistica.Operation(categoria)
	if err != nil {
		fmt.Println(err)
	}

	switch funcSelecionada {
	case "minimum":
		resultado := estatistica.MinFunc(2, 3, 3, 4, 10, 2, 4, 5)
		fmt.Println("Minimum of the list: ", resultado)
	case "average":
		resultado := estatistica.AvgFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Println("Minimum of the list: ", resultado)

	case "maximum":
		resultado := estatistica.MaxFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Println("Minimum of the list: ", resultado)

	}
}
