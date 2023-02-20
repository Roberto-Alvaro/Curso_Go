package main

import (
	"fmt"
	"reflect" //Paquete que permite a acceer a mas funciones, lo vamos a usar para TypeOf, para poder ver el tipo de dato
)

func main() {

	var cifra int = 10 //var <nombre de la variable> <Tipo de dato> = <Valor>
	// entero y flotantes/strings
	var palabra string = "guardar"
	var numero_f float64 = 2.333
	fmt.Println(cifra)
	fmt.Println(palabra)
	fmt.Println(numero_f)
	fmt.Println(reflect.TypeOf(numero_f))
	fmt.Println(reflect.TypeOf(palabra))
}
