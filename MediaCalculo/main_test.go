package main

import "testing"

func TestMediaDoAluno(t *testing.T) {
	resultado := MediaDoAluno(7, 7)
	esperado := true
	if resultado != esperado {
		t.Errorf("Média do Aluno(6, 8) == %t, esperado %t", resultado, esperado)
	}

	resultado = MediaDoAluno(7, 5)
	esperado = false
	if resultado != esperado {
		t.Errorf("Média do Aluno(6, 6) == %t, esperado %t", resultado, esperado)
	}
}
