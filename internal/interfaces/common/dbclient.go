package common

import "database/sql"

type DBClient interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	Rebind(query string) string
}
