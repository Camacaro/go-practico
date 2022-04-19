package server

import (
	m "06-webserver/middlewares"
	"06-webserver/router"
	"net/http"
)

type Server struct {
	port   string
	router *router.Router
}

func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: router.NewRouter(),
	}
}

func (s *Server) Handle(mathod string, path string, handler http.HandlerFunc) {
	_, pathExist := s.router.Rules[path]

	if !pathExist {
		s.router.Rules[path] = make(map[string]http.HandlerFunc)
	}

	s.router.Rules[path][mathod] = handler
}

//...m.Middleware is a type of Middleware slice (array dinamico)
func (s *Server) AddMiddleware(f http.HandlerFunc, middlewares ...m.Middleware) http.HandlerFunc {
	for _, middleware := range middlewares {
		f = middleware(f)
	}

	return f
}

func (s *Server) Listen() error {

	http.Handle("/", s.router)

	err := http.ListenAndServe(":"+s.port, nil)

	if err != nil {
		return err
	}

	return nil
}
