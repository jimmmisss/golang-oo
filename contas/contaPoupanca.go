package contas

import "go-oo/clientes"

type ContaPoupanca struct {
	Titular                  clientes.Titular
	Agencia, Conta, Operacao int
	saldo                    float64
}

func (c *ContaPoupanca) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorSaque
		return "Saque no valor R$"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaPoupanca) Depositar(valor float64) (string, float64) {
	if valor > 0 {
		c.saldo += valor
		return "Depósito realizado: R$", valor
	} else {
		return "Valor do depósito menor que zero", valor
	}
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
