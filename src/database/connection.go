package database

import (
	"book-api/src/config"
	"database/sql"
)

func Connection() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.ConnectionStringDb)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
