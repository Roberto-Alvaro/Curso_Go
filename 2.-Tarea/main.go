/*Crea 4 variables las cuales almacenaran datos que se le pedirán al usuario
en consola. Los datos que se le pedirán al usuario son los siguientes:

- Nombre

- Dirección

- -Edad

- -Profesión

Luego muestra todos estos datos en pantalla*/

package main

import "fmt"

func main() {
	var nombre string
	var direccion string
	var edad int
	var profesion string

	fmt.Println("Digita tu nombre")
	fmt.Scan(&nombre)
	fmt.Println("Digita tu direccion")
	fmt.Scan(&direccion)
	fmt.Println("Digita tu profesion")
	fmt.Scan(&profesion)
	fmt.Println("Digita tu edad")
	fmt.Scan(&edad)
	fmt.Println("Sus datos son: ", nombre, direccion, profesion, edad)
}
