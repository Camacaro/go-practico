package middlewares

import (
	"encoding/json"
	"net/http"
)

type Middleware func(f http.HandlerFunc) http.HandlerFunc

type MetaData interface{}

/*
	Esta notacion es para que cuando este struct sea
	decofificado por el json, se le asigne el valor
	Cuando se serialize, se le asigna el valor en minuscula

	De esta forma se puede usar el mismo nombre en el
	json y en el struct
*/
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (u *User) ToJson() ([]byte, error) {
	return json.Marshal(u)
}
