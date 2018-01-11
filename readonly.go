package gorputil

import (
	"database/sql"

	"gopkg.in/gorp.v2"
)

type readonlySqlExecutor struct {
	gorp.SqlExecutor
}

func (r *readonlySqlExecutor) Insert(list ...interface{}) error {
	panic("readonly")
}

func (r *readonlySqlExecutor) Update(list ...interface{}) (int64, error) {
	panic("readonly")
}

func (r *readonlySqlExecutor) Delete(list ...interface{}) (int64, error) {
	panic("readonly")
}

func (r *readonlySqlExecutor) Exec(query string, args ...interface{}) (sql.Result, error) {
	panic("readonly")
}

func (r *readonlySqlExecutor) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if isSelectQuery(query) {
		return r.SqlExecutor.Query(query, args)
	}
	panic("readonly")
}

func (r *readonlySqlExecutor) QueryRow(query string, args ...interface{}) *sql.Row {
	if isSelectQuery(query) {
		return r.SqlExecutor.QueryRow(query, args)
	}
	panic("readonly")
}

func ReadonlySqlExecutor(executor gorp.SqlExecutor) gorp.SqlExecutor {
	return &readonlySqlExecutor{executor}
}
