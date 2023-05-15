package database

import (
	"log"
	database "module-path/Lista-de-tarefas/database/mysql"
	"module-path/Lista-de-tarefas/tarefas"
)

func Insert() error {

	db, erro := database.Connect()
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()

	stmt, erro := db.Prepare("INSERT INTO dados(datas, hora, tarefas) VALUES (?, ?, ?)")
	if erro != nil {
		log.Fatal(erro)
		return erro
	}

	dados := tarefas.Dados()

	for _, v := range dados {
		stmt.Exec(v.Data, v.Hora, v.Tarefa)
	}
	defer stmt.Close()

	return nil
}
