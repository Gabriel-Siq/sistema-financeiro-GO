package main

import "fmt"

type ContaCorrente struct {
	titular    string
	numAgencia int
	numConta   int
	saldo      float64
}

func (conta *ContaCorrente) realizaSaque() {
	var valorSaque float64

	fmt.Print("Informe o valor que deseja sacar: ")
	fmt.Scanf("%f", &valorSaque)

	if valorSaque <= conta.saldo {
		conta.saldo -= valorSaque

		fmt.Print("O valor do saldo atual eh: ", conta.saldo)
		fmt.Println("")
	} else {
		fmt.Println("O valor desejado de saque é maior do que o existente na conta")
	}

}

func main() {

	contaTheo := ContaCorrente{}
	contaTheo.titular = "Theo"
	contaTheo.saldo = 500

	fmt.Println(contaTheo)

	contaTheo.realizaSaque()
}
