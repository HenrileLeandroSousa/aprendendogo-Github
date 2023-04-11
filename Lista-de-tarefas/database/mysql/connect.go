package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	banco := "root:123456@/bancodalista"
	db, erro := sql.Open("mysql", banco)
	if erro != nil {
		log.Fatal(erro)
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
