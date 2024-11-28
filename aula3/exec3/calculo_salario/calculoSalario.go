package calculosalario

import (
	"errors"
	"strings"
)

func CalculoSalario(minutos int, categoria string) (salario float64, err error) {
	categoria = strings.ToUpper(categoria)
	horas := minutos / 60
	switch categoria {
	case "A":
		salario = float64(3000*horas) * 1.5

	case "B":
		salario = float64(1500*horas) * 1.2

	case "C":
		salario = float64(1000 * horas)

	default:
		err = errors.New("Caso aplicado não configurado em nosso sistema")
		return 0, err

	}

	return salario, err

}
