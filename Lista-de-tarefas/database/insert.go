package database

import (
	"log"
	database "module-path/Lista-de-tarefas/database/mysql"

	_ "github.com/go-sql-driver/mysql"
)

func Insert() error {

	db, erro := database.Connect()
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()

	d := tarefas.InserirTarefas()[0]

	stmt, erro := db.Prepare("INSERT INTO dados(datas, hora, tarefas) VALUES (?, ?, ?)")
	if erro != nil {
		log.Fatal(erro)
		return erro
	}
	stmt.Exec(d.Data, d.Hora, d.Tarefa)
	defer stmt.Close()

	return nil
}
