package main

import (
	"fmt"
	"go-oo/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	contaWesley := contas.ContaPoupanca{}
	contaWesley.Depositar(1000)
	PagarBoleto(&contaWesley, 130)
	fmt.Println(contaWesley.ObterSaldo())

	contaFadia := contas.ContaCorrente{}
	contaFadia.Depositar(400)
	PagarBoleto(&contaFadia, 130)
	fmt.Println(contaFadia.ObterSaldo())

	/*clienteIsadora := clientes.Titular{Nome: "Isadora Pereira", Cpf: "789", Profissao: "Médica"}
	contaIsadora := contas.contaCorrente{clienteIsadora, 999, 666, 10000}
	fmt.Println(contaIsadora)*/

	/*contaWesley := contas.contaCorrente{Titular: "Wesley", saldo: 1300}
	contaFadia := contas.contaCorrente{Titular: "Fadia", saldo: 2000}

	status := contaWesley.Transferir(-500, &contaFadia)

	fmt.Println(status)
	fmt.Println(contaWesley)
	fmt.Println(contaFadia)

	contaWesley := contaCorrente{}
	contaWesley.titular = "Wesley"
	contaWesley.agencia = 123
	contaWesley.agencia = 321
	contaWesley.saldo = 18000.00

	fmt.Println("saldo atual R$:", contaWesley.saldo)
	fmt.Println(contaWesley.sacar(5500))
	fmt.Println("saldo atual R$:", contaWesley.saldo)
	fmt.Println(contaWesley.depositar(1000))
	fmt.Println("saldo atual R$:", contaWesley.saldo)

	status, valor := contaWesley.depositar(1000)
	fmt.Println(status, valor)*/
}
