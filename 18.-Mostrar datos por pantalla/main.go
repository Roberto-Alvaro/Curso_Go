package main

import "fmt"

func main() {
	saludo := "Hola"
	despedida := "Adios"

	fmt.Println(saludo)
	fmt.Println(despedida)

	nombre := "Roberto"
	edad := 20

	fmt.Printf("Hola %s, tu edad es: %d\n", nombre, edad) //creacion de verbos
	fmt.Printf("Hola %v, tu edad es: %v\n", nombre, edad)

	var mensaje string
	fmt.Println("ingresa un mensaje")
	fmt.Scanln(&mensaje)

	fmt.Printf("Tu mensaje es: %s", mensaje)
}
