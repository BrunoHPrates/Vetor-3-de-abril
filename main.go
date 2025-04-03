package main

import "fmt"

func main() {
	valores := []int{1, 2, 3, 4, 5}
	valores = append(valores, 6, 7, 8,)
	fmt.Println("Valores:", valores)
	fmt.Println(len(valores), cap(valores))
}