package main

import "fmt"

func main() {

	var num1 int = 20
	var num2 int = 10

	fmt.Println(num1 + num2)
	fmt.Println(num1 - num2)
	fmt.Println(num1 * num2)
	fmt.Println(num1 / num2)
	fmt.Println(num1 % num2)
	fmt.Println(num1 / num2)

	//operadores generados por las variables
	num2++ //suma jejeje
	fmt.Println(num2)
	num1--
	fmt.Println(num1)

	fmt.Println("este es un numero negativo con operador unario:", -num1) //vuelve al numero de la variable como negativo
}
