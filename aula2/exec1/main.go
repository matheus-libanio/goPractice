package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {

	leitorPalavra := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a word")

	textoInserido, err := leitorPalavra.ReadString('\n')
	if err != nil {
		errors.New("Error entering the word")
	}

	textoInserido = strings.TrimSpace(textoInserido)

	fmt.Println("number of letters: ", len(textoInserido))

	fmt.Println("The word was: ")

	for i, value := range textoInserido {
		fmt.Println("índice: ", i, "letter: ", string(value))

	}

}
