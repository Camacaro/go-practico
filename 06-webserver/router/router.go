package router

import (
	"net/http"
)

type Router struct {
	Rules map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		Rules: make(map[string]map[string]http.HandlerFunc),
	}
}

// Esto es lo que manejeja el retorno de las peticiones segun la API
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// fmt.Fprint(w, "Hello World!")
	handler, methodExist, pathExist := r.FindHandler(req.URL.Path, req.Method)

	if !pathExist {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !methodExist {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	handler(w, req)
}

func (r *Router) FindHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, pathExist := r.Rules[path]
	handler, methodExist := r.Rules[path][method]
	return handler, methodExist, pathExist
}
