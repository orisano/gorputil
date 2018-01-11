package gorputil

import (
	"database/sql"

	"gopkg.in/gorp.v2"
)

type balancedSqlExecutor struct {
	gorp.SqlExecutor
	slaves   []gorp.SqlExecutor
	balancer Balancer
}

func (b *balancedSqlExecutor) slave() gorp.SqlExecutor {
	id := b.balancer.Balance(len(b.slaves))
	return b.slaves[id]
}

func (b *balancedSqlExecutor) Select(i interface{}, query string, args ...interface{}) ([]interface{}, error) {
	return b.slave().Select(i, query, args...)
}

func (b *balancedSqlExecutor) SelectInt(query string, args ...interface{}) (int64, error) {
	return b.slave().SelectInt(query, args...)
}

func (b *balancedSqlExecutor) SelectNullInt(query string, args ...interface{}) (sql.NullInt64, error) {
	return b.slave().SelectNullInt(query, args...)
}

func (b *balancedSqlExecutor) SelectFloat(query string, args ...interface{}) (float64, error) {
	return b.slave().SelectFloat(query, args...)
}

func (b *balancedSqlExecutor) SelectNullFloat(query string, args ...interface{}) (sql.NullFloat64, error) {
	return b.slave().SelectNullFloat(query, args...)
}

func (b *balancedSqlExecutor) SelectStr(query string, args ...interface{}) (string, error) {
	return b.slave().SelectStr(query, args...)
}

func (b *balancedSqlExecutor) SelectNullStr(query string, args ...interface{}) (sql.NullString, error) {
	return b.slave().SelectNullStr(query, args...)
}

func (b *balancedSqlExecutor) SelectOne(holder interface{}, query string, args ...interface{}) error {
	return b.slave().SelectOne(holder, query, args...)
}

func (b *balancedSqlExecutor) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if isSelectQuery(query) {
		return b.slave().Query(query, args...)
	}
	return b.SqlExecutor.Query(query, args...)
}

func (b *balancedSqlExecutor) QueryRow(query string, args ...interface{}) *sql.Row {
	if isSelectQuery(query) {
		return b.slave().QueryRow(query, args...)
	}
	return b.SqlExecutor.QueryRow(query, args...)
}

func BalancedSqlExecutor(master gorp.SqlExecutor, slaves []gorp.SqlExecutor, balancer Balancer) gorp.SqlExecutor {
	return &balancedSqlExecutor{
		SqlExecutor: master,
		slaves:      slaves,
		balancer:    balancer,
	}
}
