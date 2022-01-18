package contas

import "go-oo/clientes"

type CC struct {
	Titular clientes.Titular
	Agencia int
	Conta   int
	saldo   float64
}

func (c *CC) Sacar(valorSaque float64) (string, float64) {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorSaque
		return "Saque no valor R$", valorSaque
	} else {
		return "saldo insuficiente", c.saldo
	}
}

func (c *CC) Depositar(valor float64) (string, float64) {
	if valor > 0 {
		c.saldo += valor
		return "Depósito realizado: R$", valor
	} else {
		return "Valor do depósito menor que zero", valor
	}
}

func (c *CC) Transferir(valorTransferencia float64, destino *CC) bool {
	if valorTransferencia < c.saldo && valorTransferencia > 0 {
		c.saldo -= valorTransferencia
		destino.Depositar(valorTransferencia)
		return true
	} else {
		return false
	}
}

func (c *CC) ObterSaldo() float64 {
	return c.saldo
}
