package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cal struct {
}

func (cal) operate(entrada string, operador string) {
	entradaLimpia := strings.Split(entrada, operador)
	operador1 := parsear(entradaLimpia[0])
	operador2 := parsear(entradaLimpia[1])

	switch operador {
	case "+":
		fmt.Println("El resultado es: ", operador1+operador2)
	case "-":
		fmt.Println("El resultado es: ", operador1-operador2)
	case "*":
		fmt.Println("El resultado es: ", operador1*operador2)
	case "/":
		fmt.Println("El resultado es: ", operador1/operador2)
	default:
		fmt.Println("Operador no valido")
	}
}

func parsear(entrada string) int {
	operador, err := strconv.Atoi(entrada)
	if err != nil {
		fmt.Println("Error al convertir el valor")
		return 0
	}

	return operador
}

func leerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// Refactor
func main() {
	calculadora := cal{}
	fmt.Print("Ingrese la operacion: ")
	entrada := leerEntrada()
	fmt.Println("entrada", entrada)
	fmt.Print("Ingrese el operador: ")
	operador := leerEntrada()
	fmt.Println("operador", operador)
	calculadora.operate(entrada, operador)
}

func main2() {
	// Leer el input del user
	scanner := bufio.NewScanner(os.Stdin)
	// Llamar la funcion que recibe los valores
	fmt.Print("Ingrese la operacion: ")
	scanner.Scan()

	// Obtener el valor del input
	opertation := scanner.Text()
	fmt.Println("opertation", opertation)

	// TODO: Cambiar el operador por un scanner y no en duro el valor
	operador := "/"
	// Valores sera un array de strings
	valores := strings.Split(opertation, operador)
	fmt.Println("valores", valores)
	// fmt.Println("valores[0], valores[1]", valores[0], valores[1])

	// Convertir los valores a int
	operador1, err1 := strconv.Atoi(valores[0])
	if err1 != nil {
		fmt.Println("Error al convertir el valor 1")
		return
	}

	// Convertir los valores a int
	operador2, err2 := strconv.Atoi(valores[1])
	if err2 != nil {
		fmt.Println("Error al convertir el valor 2")
		return
	}

	switch operador {
	case "+":
		fmt.Println("El resultado es: ", operador1+operador2)
	case "-":
		fmt.Println("El resultado es: ", operador1-operador2)
	case "*":
		fmt.Println("El resultado es: ", operador1*operador2)
	case "/":
		fmt.Println("El resultado es: ", operador1/operador2)
	default:
		fmt.Println("Operador no valido")
	}
}
