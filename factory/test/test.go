package test

import "fmt"

type Veiculo interface {
	PrintarQuantidadeRodas()
}

func NovoVeiculo(qtdRodas int) Veiculo {
	if qtdRodas < 0 || qtdRodas > 4 {
		return nil
	}

	if qtdRodas == 2 {
		return &moto{qtdRodas: qtdRodas}
	} else if qtdRodas == 4 {
		return &carro{qtdRodas: qtdRodas}
	}

	return nil
}

type moto struct {
	qtdRodas int
	preco    float64
}

func (c *moto) PrintarQuantidadeRodas() {
	fmt.Println(c.qtdRodas)
}

func (c *carro) PrintarQuantidadeRodas() {
	fmt.Println(c.qtdRodas)
}

type carro struct {
	qtdRodas int
	preco    float64
}

func (c *carro) DefinePreco(preco float64) {
	if preco < 5000 || preco > 1000000000 {
		fmt.Println("Nao é permitido esse valor")
		return
	}

	c.preco = preco
}

func (c *carro) DefineQtdRodas(qtdRodas int) {
	if qtdRodas < 0 || qtdRodas > 4 {
		fmt.Println("Nao é permitido esse valor")
		return
	}

	c.qtdRodas = qtdRodas
}
