package main

import "fmt"

func main() {

	figuras := make([]Area, 0)
	figuras = append(figuras, &Cuadrado{lado: 4})
	figuras = append(figuras, &Rectangulo{lado: 4, base: 6})
	figuras = append(figuras, &Triangulo{base: 4, altura: 10})

	// Calcular areas
	for _, f := range figuras {
		fmt.Println(f.CalcularArea())
	}

}

type Area interface {
	CalcularArea() int
}

// Cuadrado implementa Area de manera implicita
type Cuadrado struct {
	lado int
}

func (c *Cuadrado) CalcularArea() int {
	return c.lado * c.lado
}

// Rectangulo implementa Area de manera implicita
type Rectangulo struct {
	lado, base int
}

func (r *Rectangulo) CalcularArea() int {
	return r.lado * r.base
}

// Triangulo implementa Area de manera implicita
type Triangulo struct {
	base, altura int
}

func (t *Triangulo) CalcularArea() int {
	return (t.base * t.altura) / 2
}
