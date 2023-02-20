package main

import (
	"fmt"
	"reflect"
)

func main() {
	numeroEntero := 60
	numeroFlotante := 19.80

	var numero1 int16 = -32767 //se puede decidir el bit mediante los bits del int o float, dependiendo la memoria
	//maximo rango en el bit

	var numero2 uint16 = 10 //no acepta numeros negativos

	fmt.Println(numeroEntero, reflect.TypeOf(numeroEntero))
	fmt.Println(numeroFlotante, reflect.TypeOf(numeroFlotante))

	fmt.Println(numero1)
	fmt.Println(numero2)

	//Se usa con proyectos grandes y la memoria es limitada
}
