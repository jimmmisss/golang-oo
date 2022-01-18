package main

import (
	"fmt"
	"go-oo/contas"
)

func main() {

	contaExemplo := contas.CC{}
	contaExemplo.Depositar(-55100)
	fmt.Println(contaExemplo.ObterSaldo())

	/*clienteIsadora := clientes.Titular{Nome: "Isadora Pereira", Cpf: "789", Profissao: "Médica"}
	contaIsadora := contas.CC{clienteIsadora, 999, 666, 10000}
	fmt.Println(contaIsadora)*/

	/*contaWesley := contas.CC{Titular: "Wesley", saldo: 1300}
	contaFadia := contas.CC{Titular: "Fadia", saldo: 2000}

	status := contaWesley.Transferir(-500, &contaFadia)

	fmt.Println(status)
	fmt.Println(contaWesley)
	fmt.Println(contaFadia)

	contaWesley := CC{}
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
