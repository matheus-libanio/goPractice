package estatistica

import (
	"errors"
	"math"
)

const (
	minimum = "minimum"

	average = "average"

	maximum = "maximum"
)

func Operation(operacao int) (string, error) {
	switch operacao {
	case 1:
		return minimum, nil
	case 2:
		return average, nil
	case 3:
		return maximum, nil
	}

	return "", errors.New("invalid type")
}

func MinFunc(values ...int) int {
	var resultado float64
	for index, value := range values {
		if index < len(values)-1 {
			resultado = math.Min(float64(value), float64(values[index+1]))
		}
	}
	return int(resultado)

}
func AvgFunc(values ...int) float64 {
	var total int
	for _, value := range values {
		total += value
	}

	media := float64(total) / float64(len(values))
	return media
}
func MaxFunc(values ...int) int {

	var resultado float64
	for index, value := range values {
		if index < len(values)-1 {
			resultado = math.Max(float64(value), float64(values[index+1]))
		}
	}
	return int(resultado)

}
