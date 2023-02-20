package main

import "fmt"

func main() {
	cadena1 := "Hola Walter"
	cadena2 := "Coto"

	var numero int = 1

	fmt.Println(cadena1 + cadena2)
	fmt.Println("Hola usuario: ", cadena2)
	fmt.Println("Hola usuario: " + cadena1)
	fmt.Println(numero, "Hola usuario: "+cadena2)
}
