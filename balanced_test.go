package gorputil

import (
	"database/sql"
	"fmt"
	"testing"

	"gopkg.in/gorp.v2"
)

var errReadMaster = fmt.Errorf("read master")
var errReadOnly = fmt.Errorf("readonly")

func TestBalancedSqlExecutor_Get(t *testing.T) {
	master := &MockSqlExecutor{
		GetMock: func(i interface{}, keys ...interface{}) (interface{}, error) {
			return nil, errReadMaster
		},
	}
	slaves := []gorp.SqlExecutor{
		&MockSqlExecutor{
			GetMock: func(i interface{}, keys ...interface{}) (interface{}, error) {
				return nil, nil
			},
		},
	}
	db := BalancedSqlExecutor(master, slaves, &Sequential{})
	if _, err := db.Get(nil); err != nil {
		t.Error(err)
	}
}

func TestBalancedSqlExecutor_Select(t *testing.T) {
	master := &MockSqlExecutor{
		SelectMock: func(i interface{}, query string, args ...interface{}) ([]interface{}, error) {
			return nil, errReadMaster
		},
	}
	slaves := []gorp.SqlExecutor{
		&MockSqlExecutor{
			SelectMock: func(i interface{}, query string, args ...interface{}) ([]interface{}, error) {
				return nil, nil
			},
		},
	}
	db := BalancedSqlExecutor(master, slaves, &Sequential{})
	if _, err := db.Select(nil, ""); err != nil {
		t.Error(err)
	}
}

func TestBalancedSqlExecutor_SelectInt(t *testing.T) {
	master := &MockSqlExecutor{
		SelectIntMock: func(query string, args ...interface{}) (int64, error) {
			return 0, errReadMaster
		},
	}
	slaves := []gorp.SqlExecutor{
		&MockSqlExecutor{
			SelectIntMock: func(query string, args ...interface{}) (int64, error) {
				return 0, nil
			},
		},
	}
	db := BalancedSqlExecutor(master, slaves, &Sequential{})
	if _, err := db.SelectInt(""); err != nil {
		t.Error(err)
	}
}

func TestBalancedSqlExecutor_SelectNullInt(t *testing.T) {
	master := &MockSqlExecutor{
		SelectNullIntMock: func(query string, args ...interface{}) (sql.NullInt64, error) {
			return sql.NullInt64{}, errReadMaster
		},
	}
	slaves := []gorp.SqlExecutor{
		&MockSqlExecutor{
			SelectNullIntMock: func(query string, args ...interface{}) (sql.NullInt64, error) {
				return sql.NullInt64{}, nil
			},
		},
	}
	db := BalancedSqlExecutor(master, slaves, &Sequential{})
	if _, err := db.SelectNullInt(""); err != nil {
		t.Error(err)
	}
}

func TestBalancedSqlExecutor_SelectFloat(t *testing.T) {
	master := &MockSqlExecutor{
		SelectFloatMock: func(query string, args ...interface{}) (float64, error) {
			return 0, errReadMaster
		},
	}
	slaves := []gorp.SqlExecutor{
		&MockSqlExecutor{
			SelectFloatMock: func(query string, args ...interface{}) (float64, error) {
				return 0, nil
			},
		},
	}
	db := BalancedSqlExecutor(master, slaves, &Sequential{})
	if _, err := db.SelectFloat(""); err != nil {
		t.Error(err)
	}
}

func TestBalancedSqlExecutor_SelectNullFloat(t *testing.T) {
	master := &MockSqlExecutor{
		SelectNullFloatMock: func(query string, args ...interface{}) (sql.NullFloat64, error) {
			return sql.NullFloat64{}, errReadMaster
		},
	}
	slaves := []gorp.SqlExecutor{
		&MockSqlExecutor{
			SelectNullFloatMock: func(query string, args ...interface{}) (sql.NullFloat64, error) {
				return sql.NullFloat64{}, nil
			},
		},
	}
	db := BalancedSqlExecutor(master, slaves, &Sequential{})
	if _, err := db.SelectNullFloat(""); err != nil {
		t.Error(err)
	}
}

func TestBalancedSqlExecutor_SelectStr(t *testing.T) {
	master := &MockSqlExecutor{
		SelectStrMock: func(query string, args ...interface{}) (string, error) {
			return "", errReadMaster
		},
	}
	slaves := []gorp.SqlExecutor{
		&MockSqlExecutor{
			SelectStrMock: func(query string, args ...interface{}) (string, error) {
				return "", nil
			},
		},
	}
	db := BalancedSqlExecutor(master, slaves, &Sequential{})
	if _, err := db.SelectStr(""); err != nil {
		t.Error(err)
	}
}

func TestBalancedSqlExecutor_SelectNullStr(t *testing.T) {
	master := &MockSqlExecutor{
		SelectNullStrMock: func(query string, args ...interface{}) (sql.NullString, error) {
			return sql.NullString{}, errReadMaster
		},
	}
	slaves := []gorp.SqlExecutor{
		&MockSqlExecutor{
			SelectNullStrMock: func(query string, args ...interface{}) (sql.NullString, error) {
				return sql.NullString{}, nil
			},
		},
	}
	db := BalancedSqlExecutor(master, slaves, &Sequential{})
	if _, err := db.SelectNullStr(""); err != nil {
		t.Error(err)
	}
}

func TestBalancedSqlExecutor_SelectOne(t *testing.T) {
	master := &MockSqlExecutor{
		SelectOneMock: func(holder interface{}, query string, args ...interface{}) error {
			return errReadMaster
		},
	}
	slaves := []gorp.SqlExecutor{
		&MockSqlExecutor{
			SelectOneMock: func(holder interface{}, query string, args ...interface{}) error {
				return nil
			},
		},
	}
	db := BalancedSqlExecutor(master, slaves, &Sequential{})
	if err := db.SelectOne(nil, ""); err != nil {
		t.Error(err)
	}
}

func TestBalancedSqlExecutor_Query(t *testing.T) {
	tests := []struct {
		query         string
		master, slave gorp.SqlExecutor
	}{
		{
			query: "select 1;",
			master: &MockSqlExecutor{
				QueryMock: func(query string, args ...interface{}) (*sql.Rows, error) {
					return nil, errReadMaster
				},
			},
			slave: &MockSqlExecutor{
				QueryMock: func(query string, args ...interface{}) (*sql.Rows, error) {
					return nil, nil
				},
			},
		},
		{
			query: "SELECT * FROM foo;",
			master: &MockSqlExecutor{
				QueryMock: func(query string, args ...interface{}) (*sql.Rows, error) {
					return nil, errReadMaster
				},
			},
			slave: &MockSqlExecutor{
				QueryMock: func(query string, args ...interface{}) (*sql.Rows, error) {
					return nil, nil
				},
			},
		},
		{
			query: "DELETE FROM foo;",
			master: &MockSqlExecutor{
				QueryMock: func(query string, args ...interface{}) (*sql.Rows, error) {
					return nil, nil
				},
			},
			slave: &MockSqlExecutor{
				QueryMock: func(query string, args ...interface{}) (*sql.Rows, error) {
					return nil, errReadOnly
				},
			},
		},
		{
			query: "INSERT INTO record(id, name) VALUES(1, foo);",
			master: &MockSqlExecutor{
				QueryMock: func(query string, args ...interface{}) (*sql.Rows, error) {
					return nil, nil
				},
			},
			slave: &MockSqlExecutor{
				QueryMock: func(query string, args ...interface{}) (*sql.Rows, error) {
					return nil, errReadOnly
				},
			},
		},
		{
			query: "UPDATE foo SET name=bar;",
			master: &MockSqlExecutor{
				QueryMock: func(query string, args ...interface{}) (*sql.Rows, error) {
					return nil, nil
				},
			},
			slave: &MockSqlExecutor{
				QueryMock: func(query string, args ...interface{}) (*sql.Rows, error) {
					return nil, errReadOnly
				},
			},
		},
	}
	for _, test := range tests {
		db := BalancedSqlExecutor(test.master, []gorp.SqlExecutor{test.slave}, &Sequential{})
		if _, err := db.Query(test.query); err != nil {
			t.Error(err)
		}
	}
}

func TestBalancedSqlExecutor_QueryRow(t *testing.T) {
	tests := []struct {
		query         string
		master, slave gorp.SqlExecutor
	}{
		{
			query: "SELECT 1;",
			master: &MockSqlExecutor{
				QueryRowMock: func(query string, args ...interface{}) *sql.Row {
					return nil
				},
			},
			slave: &MockSqlExecutor{
				QueryRowMock: func(query string, args ...interface{}) *sql.Row {
					return &sql.Row{}
				},
			},
		},
		{
			query: "DELETE FROM foo;",
			master: &MockSqlExecutor{
				QueryRowMock: func(query string, args ...interface{}) *sql.Row {
					return &sql.Row{}
				},
			},
			slave: &MockSqlExecutor{
				QueryRowMock: func(query string, args ...interface{}) *sql.Row {
					return nil
				},
			},
		},
		{
			query: "INSERT INTO record(id, name) VALUES(1, foo);",
			master: &MockSqlExecutor{
				QueryRowMock: func(query string, args ...interface{}) *sql.Row {
					return &sql.Row{}
				},
			},
			slave: &MockSqlExecutor{
				QueryRowMock: func(query string, args ...interface{}) *sql.Row {
					return nil
				},
			},
		},
		{
			query: "UPDATE foo SET name=bar;",
			master: &MockSqlExecutor{
				QueryRowMock: func(query string, args ...interface{}) *sql.Row {
					return &sql.Row{}
				},
			},
			slave: &MockSqlExecutor{
				QueryRowMock: func(query string, args ...interface{}) *sql.Row {
					return nil
				},
			},
		},
	}
	for _, test := range tests {
		db := BalancedSqlExecutor(test.master, []gorp.SqlExecutor{test.slave}, &Sequential{})
		if row := db.QueryRow(test.query); row == nil {
			t.Errorf("called to unexpected executor %q", test.query)
		}
	}
}
