package main

import "fmt"

func main() {
	//uso de barra para decir que no ha cerrado el comentario
	var cadena string = "Ejemplo de \"cadena de texto, con comillas\""
	fmt.Println(cadena)

	var cadena2 string = "Ejemplo de \nsalto de linea"
	fmt.Println(cadena2)

	var cadena3 string = "Ejemplo de \t\ttabulacion"
	fmt.Println(cadena3)

	var cadena4 string = "Ejemplo de \bbarra invertida b"
	fmt.Println(cadena4)

	var cadena5 string = "Ejemplo de \fbarra invertida f"
	fmt.Println(cadena5)

	var cadena6 string = "Ejemplo de \vbarra invertida v"
	fmt.Println(cadena6)

	var cadena7 string = "Ejemplo de \rbarra invertida r"
	fmt.Println(cadena7)
}
