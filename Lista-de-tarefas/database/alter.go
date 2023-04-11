package database

import (
	"log"
	database "module-path/Lista-de-tarefas/database/mysql"

	_ "github.com/go-sql-driver/mysql"
)

func ToAlter() {

	db, erro := database.Connect()
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()

	stmt, erro := db.Prepare("UPDATE dados SET datas, hora = ( ?, ?, ?) WHERE id = (?)")
	if erro != nil {
		log.Fatal(erro)
	}
	defer stmt.Close()

}
