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
	|---Tiempo ---| = 12.6 segundos
	FB -> IG -> PZ

	Con GoRoutne (pero no se conocen entre si)
	|-T-|
	FB
	IG
	PZ

	Con Channel (Se comunican entre si)
	|-T-| = 3 segundos
	|- FB
	|- IG (Envia el ok: C <- ok)
	|- PZ (Recibe el ok: <- C)
*/

func main() {
	canal := make(chan string)

	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://youtube.com",
	}

	i := 0

	for {
		// Romper el ciclo
		if i > 2 {
			break
		}

		for _, servidor := range servidores {
			go revisarServidor(servidor, canal)
		}

		// Hacer un request a revisarServidor cada segundo
		time.Sleep(time.Second * 1)
		fmt.Println(<-canal)

		i++
	}
}

func revisarServidor(servidor string, canal chan string) {
	// Solo tengo que revisar si es un error
	_, err := http.Get(servidor)
	if err != nil {
		canal <- "Servidor caido: " + servidor
	} else {
		canal <- "Servidor funcionando: " + servidor
	}
}
