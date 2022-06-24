package postgres

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type PgDBClient struct {
	connection *sqlx.DB
	tx         *sql.Tx
}

func NewPgDbClient(conn *sqlx.DB, tx *sql.Tx) *PgDBClient {
	return &PgDBClient{
		connection: conn,
		tx:         tx,
	}
}

func (c *PgDBClient) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return c.tx.Query(query, args...)
}
func (c *PgDBClient) Exec(query string, args ...interface{}) (sql.Result, error) {
	return c.tx.Exec(query, args...)
}
func (c *PgDBClient) Rebind(query string) string {
	return c.connection.Rebind(query)
}
