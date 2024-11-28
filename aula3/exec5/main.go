package main

import (
	"fmt"

	"rsc.io/quote"
)

type abrigo struct {
}

func main() {

	fmt.Println(quote.Go())

}

func verifyAnimal(animal string) error {
	animaisValidos := []string{"dog", "cat", "hammster", "tarantula"}
	for _, animalValido := range animaisValidos {
		if animal == animalValido {
			var err error
			return err
		}
	}
	return nil
}
