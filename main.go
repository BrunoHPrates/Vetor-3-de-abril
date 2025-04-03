package main

import "fmt"

func main() {
	nomes := []string{"Bruno", "Yan", "Vinicius", "Eduardo", "Pedro"}
	fmt.Println(nomes[:2])
	fmt.Println(nomes[3:])
	rangOne := nomes[2]
	fmt.Println(rangOne)
}