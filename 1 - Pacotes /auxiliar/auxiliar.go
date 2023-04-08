package auxiliar

import "fmt"

/*
Dentro do pacote conseguimos referenciar funções do mesmo.
Já na chamada externa podemos somente chamar funções públicas se a mesma começar com sintaxe de letra maiúscula.
*/
func Escrever() {
	fmt.Println("Escrevendo do pacote auxiliar")
	escrever2()
}
