package main

type abrigo struct {
}

func main() {

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
