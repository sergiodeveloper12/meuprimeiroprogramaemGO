package main // geralmente se usa "main" pra programa executável

import "fmt"

func main() {
	fmt.Print("Digite seu nome: ") // Print sem quebrar linha
	var nome string                // a variável precisa ser "nome", igual no Scanln
	fmt.Scanln(&nome)

	fmt.Println("Bem vindo", nome) // corrigido: fmt.Println, não fmt,Println

	fmt.Println("sua idade é: ")
	var idade int
	fmt.Scanln(&idade)
	fmt.Println("Você tem", idade, "anos")
	fmt.Print("Digite seu peso:")
	var peso float64
	fmt.Scanln(&peso)
	fmt.Println("Seu peso é:", peso, "kg")
	fmt.Print("Digite sua altura:")
	var altura float64
	fmt.Scanln(&altura)
	fmt.Println("Sua altura é:", altura, "m")
	imc := peso / (altura * altura)
	fmt.Printf("Seu IMC é: %.2f\n", imc) // corrigido: fmt.Printf para formatar a saída
}
