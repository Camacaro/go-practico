package main

import (
	"fmt"
	"net/http"
	"time"
)

/*
	Verificar si los servidores estan funcionando

	Es una aplicacion que estara verificando que los servidores esten funcionando

	Con GoRoutne vamos a ejecutar estos servidores de manera simultanea
	pero vamos aplicar Channel para que se conozcan entre si esas GoRoutne

	Normal
	|---Tiempo ---| = 12 segundos
	FB -> IG -> PZ

	Con GoRoutne (pero no se conocen entre si)
	|-T-|
	FB
	IG
	PZ

	Con Channel (Se comunican entre si)
	|-T-|
	|- FB
	|- IG (Envia el ok: C <- ok)
	|- PZ (Recibe el ok: <- C)
*/

func main() {

	// Hora actual
	inicio := time.Now()

	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://youtube.com",
	}

	for _, servidor := range servidores {
		revisarServidor(servidor)
	}

	// Tiempo que tardo en ejecutarse
	tiempoPaso := time.Since(inicio)
	fmt.Println("Tiempo de ejecucion:", tiempoPaso)
}

func revisarServidor(servidor string) {
	// Solo tengo que revisar si es un error
	_, err := http.Get(servidor)
	if err != nil {
		fmt.Println("Servidor caido:", servidor)
	} else {
		fmt.Println("Servidor funcionando:", servidor)
	}
}
