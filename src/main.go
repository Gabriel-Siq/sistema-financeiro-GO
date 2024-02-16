package main

import "fmt"

type ContaCorrente struct {
	titular    string
	numAgencia int
	numConta   int
	saldo      float64
}

func (c *ContaCorrente) realizaIntroducao(contaTheo ContaCorrente) {
	escolha := 0
	fmt.Println("1 - Sacar")
	fmt.Println("2 - Depositar")
	fmt.Println("3 - Sair")
	fmt.Print("Informe a operação que deseja realizar: ")
	fmt.Scanf("%d", &escolha)
	fmt.Println("-------------------------")

	switch escolha {
	case 1:
		contaTheo.realizaSaque()
	case 2:
		contaTheo.realizaDeposito()
	case 3:
		fmt.Println("Sem problemas, até a próxima :)")
	}
}

func (conta *ContaCorrente) realizaSaque() {
	var valorSaque float64

	fmt.Println("Seu saldo atual é: ", conta.saldo)
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

func (conta *ContaCorrente) realizaDeposito() {
	var valorDeposito float64

	fmt.Print("Seu saldo atual é: ", conta.saldo)
	fmt.Println("")

	fmt.Print("Informe o valor que deseja depositar: ")
	fmt.Scanf("%f", &valorDeposito)

	if valorDeposito > 0 {
		conta.saldo += valorDeposito

		fmt.Print("O valor do saldo atual eh: ", conta.saldo)
		fmt.Println("")
	} else {
		fmt.Println("O valor inserido é indisponível, peço desculpas :(")
	}
}

func indentificaUsuario() (string, int, int, float64) {
	nome := ""
	numAgencia := 0
	numConta := 0
	saldo := 12.

	fmt.Println("Precisaremos de seus dados para acessar sua conta.")
	fmt.Print("Informe seu nome: ")
	fmt.Scanf("%s", &nome)
	fmt.Print("Informe o número da sua agência: ")
	fmt.Scanf("%d", &numAgencia)
	fmt.Print("Informe o número da sua conta: ")
	fmt.Scanf("%d", &numConta)
	fmt.Print("Informe seu saldo: ")
	fmt.Scanf("%f", &saldo)
	fmt.Println("-------------------------")

	return nome, numAgencia, numConta, saldo
}

func main() {
	nome, numAgencia, numConta, saldo := indentificaUsuario()

	contaTheo := ContaCorrente{}
	contaTheo.titular = nome
	contaTheo.numAgencia = numAgencia
	contaTheo.numConta = numConta
	contaTheo.saldo = saldo

	contaTheo.realizaIntroducao(contaTheo)

}
