package main

import "fmt"

func main() {
	fmt.Print("Informe o texto que você deseja inverter: ")
	var text string
	fmt.Scan(&text)

	var reverseText = reverseString(text)
	fmt.Println("Aqui está seu texto invertido!")
	fmt.Println(reverseText)
}

func reverseString(text string) string {
	charSlice := []byte(text)
    
	size := len(charSlice)

	for i := 0; i < size/2; i++ {

		j := size - i - 1	

		charSlice[i], charSlice[j] = charSlice[j], charSlice[i]
	}

	return string(charSlice)
}
