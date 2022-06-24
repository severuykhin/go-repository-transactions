package postgres

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type PgDBClient struct {
	connection *sqlx.DB
}

func NewPgDbClient(conn *sqlx.DB) *PgDBClient {
	return &PgDBClient{
		connection: conn,
	}
}

func (c *PgDBClient) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return c.connection.Query(query, args...)
}
func (c *PgDBClient) Exec(query string, args ...interface{}) (sql.Result, error) {
	return c.connection.Exec(query, args...)
}
func (c *PgDBClient) Rebind(query string) string {
	return c.connection.Rebind(query)
}
func (c *PgDBClient) QueryRow(query string, args ...interface{}) *sql.Row {
	return c.connection.QueryRow(query, args...)
}
