package main

import (
	"fmt"
 	"strings"
)

func main() {
	var texto string
	fmt.Print("Digite uma Palavra: ")
	fmt.Scanln(&texto)

	texto = strings.ToLower(strings.ReplaceAll(texto, " ", ""))

	palindromo := true
	for i := 0; i < len(texto)/2; i++ {
		if texto[i] != texto[len(texto)-1-i] {
			palindromo = false
			break
		}
	}

	if palindromo {
		fmt.Printf("%s é um palíndromo\n", texto)
	} else {
		fmt.Printf("%s não é um palíndromo\n", texto)
	}
}