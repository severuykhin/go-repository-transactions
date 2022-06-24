package postgres

import (
	"fmt"
	"time"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

func NewPostgresConnection(connString string) *sqlx.DB {
	for {
		if db, err := tryCreatePostgresConnection(connString); err == nil {
			return db
		}

		time.Sleep(time.Second)
	}
}

func tryCreatePostgresConnection(connStr string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", connStr)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return nil, err
	}

	return db, nil
}
