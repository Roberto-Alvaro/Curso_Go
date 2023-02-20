package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hola mundo")
	var cadena1 string = "Roberto Alvaro"
	fmt.Println(reflect.TypeOf(cadena1))
	var caracter = 'W'
	fmt.Println(reflect.TypeOf(caracter))

}
