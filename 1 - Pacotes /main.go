package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail" // Import external packate example
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("devbook@gmail.com")
	fmt.Println(erro)
}
