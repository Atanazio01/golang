package main

import (
	"errors"
	"fmt"
)

func main() {
	// INTEIROS
	var numero int64 = 1000000000000000000
	println(numero)

	var numero2 uint32 = 100000000 // unsygned int
	println(numero2)

	//alias
	// INT32 = RUNE
	var numero3 rune = 123456
	println(numero3)

	// BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	//REAIS
	// não pode ser declarado apenas como float: var numeroReal1 float = 123.45
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 12364647474984.45684
	fmt.Println(numeroReal2)

	numeroReal3 := 123456.789
	fmt.Println(numeroReal3)

	//STRINGS
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'A'
	fmt.Println(char)

	//variavel declarada sem valor pega o valor inicial do tipo
	var texto int
	fmt.Println(texto)

	// BOOLEAN
	var booleano1 bool = true
	fmt.Println(booleano1)

	//ERRO
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
