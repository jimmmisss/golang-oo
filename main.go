package main

import "fmt"

type ContaCoarrente struct {
	titular string
	agencia int
	conta   int
	saldo   float64
}

func (c *ContaCoarrente) sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCoarrente) depositar(valor float64) (string, float64) {
	if valor > 0 {
		c.saldo += valor
		return "Depósito realizado com sucesso", c.saldo
	} else {
		return "Valor do depósito menor que zero", c.saldo
	}
}

func main() {
	contaWesley := ContaCoarrente{}
	contaWesley.titular = "Wesley"
	contaWesley.agencia = 123
	contaWesley.agencia = 321
	contaWesley.saldo = 18000.00

	fmt.Println(contaWesley.saldo)
	fmt.Println(contaWesley.sacar(5500))
	fmt.Println(contaWesley.saldo)
	fmt.Println(contaWesley.depositar(1000))
	fmt.Println(contaWesley.saldo)
}
