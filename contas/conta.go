package contas

import "go-oo/clientes"

type CC struct {
	Titular clientes.Titular
	Agencia int
	Conta   int
	Saldo   float64
}

func (c *CC) sacar(valorSaque float64) (string, float64) {
	podeSacar := valorSaque > 0 && valorSaque <= c.Saldo

	if podeSacar {
		c.Saldo -= valorSaque
		return "Saque no valor R$", valorSaque
	} else {
		return "Saldo insuficiente", c.Saldo
	}
}

func (c *CC) depositar(valor float64) (string, float64) {
	if valor > 0 {
		c.Saldo += valor
		return "Depósito realizado: R$", valor
	} else {
		return "Valor do depósito menor que zero", valor
	}
}

func (c *CC) Transferir(valorTransferencia float64, destino *CC) bool {
	if valorTransferencia < c.Saldo && valorTransferencia > 0 {
		c.Saldo -= valorTransferencia
		destino.depositar(valorTransferencia)
		return true
	} else {
		return false
	}
}
