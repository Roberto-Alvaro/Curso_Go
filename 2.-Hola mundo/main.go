package main

import "fmt"

//este paquete es el mas utilizado, lo vamos a usar para mostrar cosas en pantalla, para el hola mundo
//paquete conjunto de instrucciones
func main() {
	fmt.Println("Hola mundo en Go :o") // se usa fmt.<funcion>(<sentencia>)
	fmt.Println("jejejejjejeje")
}

// instalar dependencias de GO
// go env -w GO111MODULE=off
// go get -v golang.org/x/tools/gopls
// Ejecutar un paquete o archivo .go
// go run .\main.go   --> en la direccion del programa
