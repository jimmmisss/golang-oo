package contas

import "go-oo/clientes"

type ContaCorrente struct {
	Titular        clientes.Titular
	Agencia, Conta int
	saldo          float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorSaque
		return "Saque no valor R$"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valor float64) (string, float64) {
	if valor > 0 {
		c.saldo += valor
		return "Depósito realizado: R$", valor
	} else {
		return "Valor do depósito menor que zero", valor
	}
}

func (c *ContaCorrente) Transferir(valorTransferencia float64, destino *ContaCorrente) bool {
	if valorTransferencia < c.saldo && valorTransferencia > 0 {
		c.saldo -= valorTransferencia
		destino.Depositar(valorTransferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
