package main

import "fmt"

type retangulo struct {
	altura      float32
	comprimento float32
}

func Exerc2() []float32 {
	retangulos := []retangulo{
		{altura: 12.32, comprimento: 20},
		{altura: 5, comprimento: 12.45},
		{altura: 100.12, comprimento: 200},
	}

	areas := make([]float32, len(retangulos))

	for i, v := range retangulos {
		areas[i] = v.altura * v.comprimento
	}

	return areas
}

func main() {
	areas := Exerc2()

	if len(areas) == 0 {
		fmt.Println("Não há dados passados para cálculos")
	}

	for i, area := range areas {
		fmt.Printf("Área do retângulo %d: %.2f\n", i+1, area)
	}
}
