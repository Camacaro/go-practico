package main

import "fmt"

func main() {
	mapa1 := make(map[string]int)

	mapa1["a"] = 8

	fmt.Println("Map", mapa1)
	fmt.Println("Value a", mapa1["a"])
}
