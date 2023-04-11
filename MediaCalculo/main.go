package main

import "fmt"

func MediaDoAluno(a, b float64) bool {
	resultado := (a + b) / 2
	return resultado >= 7

}

func main() {
	if MediaDoAluno(6.2, 7.8) {
		fmt.Println("O aluno está aprovado")
	} else {
		fmt.Println("O aluno está reprovado")
	}
}
