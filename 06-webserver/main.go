package main

/*
	NewServer esta en otro archivo, pero como su paquete
	es main, se puede importar directamente.

*/

import (
	h "06-webserver/handlers"
	m "06-webserver/middlewares"
	s "06-webserver/server"
)

func main() {

	server := s.NewServer("4000")
	server.Handle("GET", "/", h.HandleRoot)
	/*
		La ultima se ejecuta de primero
		CheckAuth() -> Logging() -> HandleRoot()
		CheckAuth() -> Logging() -> HandleHome()
	*/
	server.Handle("POST", "/api", server.AddMiddleware(h.HandleHome, m.Logging(), m.CheckAuth()))
	server.Handle("POST", "/api/create", h.PostRequest)
	server.Listen()
}
