package model

import (
	"database/sql"
)

type Models struct {
}

func NewModels(db *sql.DB) Models {
	// infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	return Models{}
}