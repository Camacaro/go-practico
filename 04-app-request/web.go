package main

import (
	"fmt"
	"io"
	"net/http"
)

type escritorWeb struct{}

func (escritorWeb) Write(p []byte) (int, error) {
	fmt.Println("Escribiendo...")
	// pasar de bytes a string
	fmt.Println(string(p))
	return len(p), nil
}

/*
	https://pkg.go.dev/io#Copy
	https://pkg.go.dev/io#Writer
	https://pkg.go.dev/io#ReadCloser

*/
func main() {
	// Obtener la pagina de google
	respuesta, err := http.Get("http://www.google.com")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	e := escritorWeb{}
	// https://pkg.go.dev/io#Copy
	io.Copy(e, respuesta.Body)

	// fmt.Println(respuesta.Body)
}
