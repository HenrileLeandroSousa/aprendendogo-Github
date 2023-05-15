package main

import (
	"fmt"
	"module-path/Lista-de-tarefas/database"
	"module-path/Lista-de-tarefas/tarefas"
)

func main() {
	database.Insert()
	fmt.Println(tarefas.Dados())
}
