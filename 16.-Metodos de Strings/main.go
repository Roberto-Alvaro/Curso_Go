package main

import (
	"fmt"
	"strings"
)

func main() {
	var cadena1 string = "CADENA EN MAYUSCULA"
	fmt.Println(cadena1)
	cadena1 = strings.ToLower(cadena1) //mandamos a llamar al metodo y lo aplicamos en la variable
	//Todos los caracteres el toLower los hace minusculas
	fmt.Println(cadena1)
	cadena1 = strings.ToUpper(cadena1) //Todo lo hace mayusculas
	fmt.Println(cadena1)

	cadena1 = strings.Replace(cadena1, "A", "-", 2) //Remplasa las "a"
	fmt.Println(cadena1)

	cadena1 = strings.ReplaceAll(cadena1, "A", ".") //Remplasa todo
	fmt.Println(cadena1)

}
