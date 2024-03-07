package main

import "fmt"

func main() {
	indice := 13
	soma, k := 0, 0

	for k < indice {
		k += 1
		soma += k
	}

	fmt.Println("Resultado da Soma: ", soma)
}
