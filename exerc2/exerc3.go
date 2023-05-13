package main

import (
	"fmt"
	"strings"
)

func sliceString(str []string, caracter string) []string {
	palavras := make([]string, len(str))
	for _, v := range str {
		if strings.HasPrefix(strings.ToLower(v), strings.ToLower(caracter)) {
			palavras = append(palavras, v)
		}

	}

	return palavras
}

func main() {
	str := []string{"casa", "cachorro", "garfo", "gato", "jeba", "Janeiro"}
	caracter := "c"

	v := sliceString(str, caracter)
	if len(v) == 0 {
		fmt.Println("Não há dados")
	}

	for _, c := range v {
		fmt.Println(c)
	}

}
