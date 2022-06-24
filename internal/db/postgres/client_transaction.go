package postgres

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type PgDBTransactionClient struct {
	connection *sqlx.DB
	tx         *sql.Tx
}

func NewPgDbTransactionClient(conn *sqlx.DB, tx *sql.Tx) *PgDBTransactionClient {
	return &PgDBTransactionClient{
		connection: conn,
		tx:         tx,
	}
}

func (c *PgDBTransactionClient) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return c.tx.Query(query, args...)
}
func (c *PgDBTransactionClient) QueryRow(query string, args ...interface{}) *sql.Row {
	return c.tx.QueryRow(query, args...)
}
func (c *PgDBTransactionClient) Exec(query string, args ...interface{}) (sql.Result, error) {
	return c.tx.Exec(query, args...)
}
func (c *PgDBTransactionClient) Rebind(query string) string {
	return c.connection.Rebind(query)
}
