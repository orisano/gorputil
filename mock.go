package gorputil

import (
	"database/sql"
	"gopkg.in/gorp.v2"
	"context"
)

type MockSqlExecutor struct {
	WithContextMock     func(ctx context.Context) gorp.SqlExecutor
	GetMock             func(i interface{}, keys ...interface{}) (interface{}, error)
	InsertMock          func(list ...interface{}) error
	UpdateMock          func(list ...interface{}) (int64, error)
	DeleteMock          func(list ...interface{}) (int64, error)
	ExecMock            func(query string, args ...interface{}) (sql.Result, error)
	SelectMock          func(i interface{}, query string, args ...interface{}) ([]interface{}, error)
	SelectIntMock       func(query string, args ...interface{}) (int64, error)
	SelectNullIntMock   func(query string, args ...interface{}) (sql.NullInt64, error)
	SelectFloatMock     func(query string, args ...interface{}) (float64, error)
	SelectNullFloatMock func(query string, args ...interface{}) (sql.NullFloat64, error)
	SelectStrMock       func(query string, args ...interface{}) (string, error)
	SelectNullStrMock   func(query string, args ...interface{}) (sql.NullString, error)
	SelectOneMock       func(holder interface{}, query string, args ...interface{}) error
	QueryMock           func(query string, args ...interface{}) (*sql.Rows, error)
	QueryRowMock        func(query string, args ...interface{}) *sql.Row
}

func (m *MockSqlExecutor) WithContext(ctx context.Context) gorp.SqlExecutor {
	return m.WithContextMock(ctx)
}

func (m *MockSqlExecutor) Get(i interface{}, keys ...interface{}) (interface{}, error) {
	return m.GetMock(i, keys...)
}

func (m *MockSqlExecutor) Insert(list ...interface{}) error {
	return m.InsertMock(list...)
}

func (m *MockSqlExecutor) Update(list ...interface{}) (int64, error) {
	return m.UpdateMock(list...)
}

func (m *MockSqlExecutor) Delete(list ...interface{}) (int64, error) {
	return m.DeleteMock(list...)
}

func (m *MockSqlExecutor) Exec(query string, args ...interface{}) (sql.Result, error) {
	return m.ExecMock(query, args...)
}

func (m *MockSqlExecutor) Select(i interface{}, query string, args ...interface{}) ([]interface{}, error) {
	return m.SelectMock(i, query, args...)
}

func (m *MockSqlExecutor) SelectInt(query string, args ...interface{}) (int64, error) {
	return m.SelectIntMock(query, args...)
}

func (m *MockSqlExecutor) SelectNullInt(query string, args ...interface{}) (sql.NullInt64, error) {
	return m.SelectNullIntMock(query, args...)
}

func (m *MockSqlExecutor) SelectFloat(query string, args ...interface{}) (float64, error) {
	return m.SelectFloatMock(query, args...)
}

func (m *MockSqlExecutor) SelectNullFloat(query string, args ...interface{}) (sql.NullFloat64, error) {
	return m.SelectNullFloatMock(query, args...)
}

func (m *MockSqlExecutor) SelectStr(query string, args ...interface{}) (string, error) {
	return m.SelectStrMock(query, args...)
}

func (m *MockSqlExecutor) SelectNullStr(query string, args ...interface{}) (sql.NullString, error) {
	return m.SelectNullStrMock(query, args...)
}

func (m *MockSqlExecutor) SelectOne(holder interface{}, query string, args ...interface{}) error {
	return m.SelectOneMock(holder, query, args...)
}

func (m *MockSqlExecutor) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return m.QueryMock(query, args...)
}

func (m *MockSqlExecutor) QueryRow(query string, args ...interface{}) *sql.Row {
	return m.QueryRowMock(query, args...)
}
