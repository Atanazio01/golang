package main

import "fmt"

func main() {
	// declaração explicita
	var variavel1 string = "Variável 1"

	// inferência de tipo / declaração implicita
	variavel2 := "Variável 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	// explicita com várias variaveis
	var (
		variavel3 string = "Variável 3"
		variavel4 string = "Variável 4"
	)

	fmt.Println(variavel3, variavel4)

	//implicita com várias variaveis
	variavel5, variavel6 := "Variável 5", "Variável 6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	// inverter os valores de variaveis sem precisar de uma var auxiliar
	variavel5, variavel6 = variavel6, variavel5

	fmt.Println(variavel5, variavel6)
}
