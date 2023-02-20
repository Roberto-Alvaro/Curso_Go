package main

import "fmt"

func main() {
	var cadena string = "Ejemplo de cadena de texto con debanado"
	fmt.Println(cadena[0]) //codigo Ascii de una letra en este caso "E"
	fmt.Println(cadena[0:15])

	var cadena2 string = "Ejemplo de cadena de texto con debanado"
	//fmt.Println(cadena[0]) //codigo Ascii de una letra en este caso "E"
	fmt.Println("un indice mas: ", cadena2[0:16])

	fmt.Println(cadena[:20])
	fmt.Println(cadena[20:])
}
