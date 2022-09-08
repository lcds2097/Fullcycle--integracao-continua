package main

import "fmt"

func main() {
	fmt.Println(Somar(10, 10))
}

func Somar(a int, b int) int {
	return a + b
}

func Subtrair(a int, b int) int {
	return a - b
}

func Multiplicar(a int, b int) int {
	return a * b
}

func Dividir(a int, b int) int {
	return a / b
}