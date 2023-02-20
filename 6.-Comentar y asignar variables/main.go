package main

import "fmt" //Este es un paquete que permite mostrar mensajes en pantalla

//forma no1 de declarar variables
var num int

//forma no2
var texto string = "simon"

func main() {
	fmt.Println(num)
	fmt.Println(texto)

	//forma no.3
	var texto2 string
	fmt.Println(texto2)

	//forma no.4
	var num2 int
	fmt.Println(num2)

	//forma no.5
	variable := 'w' //devuelve el codigo asquii
	variable2 := "Hola soy una variable"

	fmt.Println(variable)
	fmt.Println(variable2)

	//forma no6
	var v1, v2, v3 int
	println(v1)
	println(v2)
	println(v3)

	//froma no7
	variable4, variable5, variable6 := 10, 20, 30

	println(variable4)
	println(variable5)
	println(variable6)

}
