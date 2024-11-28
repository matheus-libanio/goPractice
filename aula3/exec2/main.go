package main

import (
	"fmt"

	"github.com/matheus-libanio/goPractice/exec2/media"
)

func main() {

	media := media.MediaNotas(4, 5, 6, 7, 12, 3, 4, 2, 5.5, 2.6, 7.8, 10)

	fmt.Printf("Média dos valores igual a: %.2f", media)

}
