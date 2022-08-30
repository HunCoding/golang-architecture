package main

import (
	"fmt"

	"github.com/HunCoding/golang-architecture/factory/test"
)

func main() {
	c := test.NovoVeiculo(20)
	fmt.Println(c)
}
