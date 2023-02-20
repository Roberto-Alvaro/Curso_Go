package main

func main() {

	/*Crea 4 variables las
	cuales 1 será decimal y las otras 3 enteras, la primera variable se
	restara con la segunda, el resultado de la resta se multiplicara con
	la tercer variable y se dividirá entre la cuarta variable.*/

	var numero1 float32 = 10.5
	var numero2 int = 10
	var numero3 int = 20
	var numero4 int = 30
	resta1 := numero1 - float32(numero2)
	multiplicacion := resta1 * float32(numero3)
	division := multiplicacion / float32(numero4)
	print(division)
}
