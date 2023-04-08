package main

import (
	"fmt"
	"module/auxiliar"

	"github.com/badoux/checkmail" // Import external packate example
)

func main() {
	fmt.Println("writing from the main file")
	auxiliar.Write()

	erro_r := checkmail.ValidateFormat("devbook@gmail.com")
	fmt.Println(erro_r)
}
