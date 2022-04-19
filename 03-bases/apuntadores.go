package main

import "fmt"

func main() {
	x := 25
	fmt.Println("Valor de X", x)
	fmt.Println("Posicion de X", &x)

	cambiarValor(x)
	fmt.Println("Valor de X al cambiar", x)

	y := &x
	fmt.Println("Valor de Y", y)
	fmt.Println("Valor de *Y", *y)
}

func cambiarValor(a int) {
	fmt.Println("Posicion de X", &a)
	a = 50
}
