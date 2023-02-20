package main

import "fmt"

func main() {
	var variable int = 1
	fmt.Println(variable)
	//cambiando el valor de variable
	variable = 25
	fmt.Println("El cambio fue hecho: ", variable)

	variable = 100
	fmt.Println("El cambio fue hecho: ", variable)

	//Constante valor que nunca va a cambiar
	//const<nombre de la variable> = <valor>
	const constante = 9.8 //fuerza gravitacional
	fmt.Print("La constante de fuerza gravitacional es: ", constante)
}
