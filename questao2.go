package main

import "fmt"

func main() {
	fmt.Print("Informe um número inteiro válido: ")
	var number int
	fmt.Scan(&number)

	if isFibonacci(number) {
		fmt.Println("O número informado faz parte da sequência de Fibonacci!")
	} else {
		fmt.Println("O número informado não faz parte da sequência de Fibonacci!")
	}
}

func isFibonacci(number int) bool {
	if number == 0 || number == 1 {
		return true
	}

	anterior, atual := 0, 1

	for atual <= number {
		if atual == number {
			return true
		}
		anterior, atual = atual, anterior+atual
	}

	return false
}
