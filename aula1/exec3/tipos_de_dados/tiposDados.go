package tiposdedados

import (
	"errors"
	"fmt"
)

var (
	teste string = "String de teste"
)

func RespostaDosExercicios() (int, error) {
	var lastName string = "Smith"
	var age int = 35
	booleano := true
	firstName := "Mary"

	fmt.Println("As variaveis corretas são lastName e firstName", lastName, firstName)

	fmt.Println("idade: ", age, "\nboleano: ", booleano)

	if age < 40 {
		return age, errors.New("idade invalida")
	}

	return 45, nil

}
