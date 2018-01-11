package gorputil

import (
	"database/sql"
	"testing"
)

func passReadonly() {
	if err := recover(); err != "readonly" {
		panic(err)
	}
}

func TestReadonlySqlExecutor(t *testing.T) {
	t.Run("Get", func(t *testing.T) {
		db := ReadonlySqlExecutor(&MockSqlExecutor{
			GetMock: func(i interface{}, keys ...interface{}) (interface{}, error) {
				return "foo", nil
			},
		})
		if got, _ := db.Get(""); got != "foo" {
			t.Errorf("unexpected value. expected: %v, but got: %v", "foo", got)
		}
	})
	t.Run("Insert", func(t *testing.T) {
		defer passReadonly()
		db := ReadonlySqlExecutor(&MockSqlExecutor{
			InsertMock: func(list ...interface{}) error {
				return nil
			},
		})
		db.Insert()
		t.Errorf("succeeded to insert")
	})
	t.Run("Update", func(t *testing.T) {
		defer passReadonly()
		db := ReadonlySqlExecutor(&MockSqlExecutor{
			UpdateMock: func(list ...interface{}) (int64, error) {
				return 1, nil
			},
		})
		db.Update()
		t.Errorf("succeeded to update")
	})
	t.Run("Delete", func(t *testing.T) {
		defer passReadonly()
		db := ReadonlySqlExecutor(&MockSqlExecutor{
			DeleteMock: func(list ...interface{}) (int64, error) {
				return 1, nil
			},
		})
		db.Delete()
		t.Errorf("succeeded to delete")
	})
	t.Run("Exec", func(t *testing.T) {
		defer passReadonly()
		db := ReadonlySqlExecutor(&MockSqlExecutor{
			ExecMock: func(query string, args ...interface{}) (sql.Result, error) {
				return nil, nil
			},
		})
		db.Exec("")
		t.Errorf("succeeded to exec")
	})
	t.Run("Select", func(t *testing.T) {
		expected := []interface{}{"foo", "bar"}
		db := ReadonlySqlExecutor(&MockSqlExecutor{
			SelectMock: func(i interface{}, query string, args ...interface{}) ([]interface{}, error) {
				return expected, nil
			},
		})
		equals := func(a, b []interface{}) bool {
			if len(a) != len(b) {
				return false
			}
			for i := range a {
				if a[i] != b[i] {
					return false
				}
			}
			return true
		}
		var dummy int
		if got, _ := db.Select(&dummy, ""); !equals(got, expected) {
			t.Errorf("unexpected result. expected: %v, but got: %v", expected, got)
		}
	})
	t.Run("SelectInt", func(t *testing.T) {
		db := ReadonlySqlExecutor(&MockSqlExecutor{
			SelectIntMock: func(query string, args ...interface{}) (int64, error) {
				return 1, nil
			},
		})
		if got, _ := db.SelectInt(""); got != 1 {
			t.Errorf("unexpected result. expected: 1, but got: %v", got)
		}
	})
	t.Run("SelectNullInt", func(t *testing.T) {
		db := ReadonlySqlExecutor(&MockSqlExecutor{
			SelectNullIntMock: func(query string, args ...interface{}) (sql.NullInt64, error) {
				return sql.NullInt64{Int64: 1, Valid: true}, nil
			},
		})
		if got, _ := db.SelectNullInt(""); got.Int64 != 1 {
			t.Errorf("unexpected result. expected: 1, but got: %v", got.Int64)
		}
	})
	t.Run("SelectFloat", func(t *testing.T) {
		db := ReadonlySqlExecutor(&MockSqlExecutor{
			SelectFloatMock: func(query string, args ...interface{}) (float64, error) {
				return 1.0, nil
			},
		})
		if got, _ := db.SelectFloat(""); got != 1.0 {
			t.Errorf("unexpected result. expected: 1.0, but got: %v", got)
		}
	})
	t.Run("SelectNullFloat", func(t *testing.T) {
		db := ReadonlySqlExecutor(&MockSqlExecutor{
			SelectNullFloatMock: func(query string, args ...interface{}) (sql.NullFloat64, error) {
				return sql.NullFloat64{Float64: 1.0, Valid: true}, nil
			},
		})
		if got, _ := db.SelectNullFloat(""); got.Float64 != 1.0 {
			t.Errorf("unexpected result. expected: 1.0, but got: %v", got.Float64)
		}
	})
	t.Run("SelectStr", func(t *testing.T) {
		db := ReadonlySqlExecutor(&MockSqlExecutor{
			SelectStrMock: func(query string, args ...interface{}) (string, error) {
				return "foo", nil
			},
		})
		if got, _ := db.SelectStr(""); got != "foo" {
			t.Errorf("unexpected result. expected: %q, but got: %q", "foo", got)
		}
	})
	t.Run("SelectNullStr", func(t *testing.T) {
		db := ReadonlySqlExecutor(&MockSqlExecutor{
			SelectNullStrMock: func(query string, args ...interface{}) (sql.NullString, error) {
				return sql.NullString{String: "foo", Valid: true}, nil
			},
		})
		if got, _ := db.SelectNullStr(""); got.String != "foo" {
			t.Errorf("unexpected result. expected: %q, but got: %q", "foo", got.String)
		}
	})
	t.Run("SelectOne", func(t *testing.T) {
		db := ReadonlySqlExecutor(&MockSqlExecutor{
			SelectOneMock: func(holder interface{}, query string, args ...interface{}) error {
				x := holder.(*int)
				*x = 1
				return nil
			},
		})
		var got int
		if db.SelectOne(&got, ""); got != 1 {
			t.Errorf("unexpected result. expected: 1, but got: %v", got)
		}
	})
	t.Run("Query", func(t *testing.T) {
		tests := []struct{
			query string
			isSelect bool
		} {
			{
				query: "SELECT 1;",
				isSelect: true,
			},
			{
				query: "select 1;",
				isSelect: true,
			},
			{
				query: "delete from foo;",
				isSelect: false,
			},
			{
				query: "insert into record(id, name) values(1, foo);",
				isSelect: false,
			},
			{
				query: "update foo set name=bar;",
				isSelect: false,
			},
		}
		for _, test := range tests {
			t.Run("Case", func(t *testing.T) {
				if !test.isSelect {
					defer passReadonly()
				}
				db := ReadonlySqlExecutor(&MockSqlExecutor{
					QueryMock: func(query string, args ...interface{}) (*sql.Rows, error) {
						return nil, nil
					},
				})
				db.Query(test.query)
				if !test.isSelect {
					t.Errorf("executed to query: %q", test.query)
				}
			})
		}
	})
	t.Run("QueryRow", func(t *testing.T) {
		tests := []struct{
			query string
			isSelect bool
		} {
			{
				query: "SELECT 1;",
				isSelect: true,
			},
			{
				query: "select 1;",
				isSelect: true,
			},
			{
				query: "delete from foo;",
				isSelect: false,
			},
			{
				query: "insert into record(id, name) values(1, foo);",
				isSelect: false,
			},
			{
				query: "update foo set name=bar;",
				isSelect: false,
			},
		}
		for _, test := range tests {
			t.Run("Case", func(t *testing.T) {
				if !test.isSelect {
					defer passReadonly()
				}
				db := ReadonlySqlExecutor(&MockSqlExecutor{
					QueryRowMock: func(query string, args ...interface{}) *sql.Row {
						return nil
					},
				})
				db.QueryRow(test.query)
				if !test.isSelect {
					t.Errorf("executed to query: %q", test.query)
				}
			})
		}
	})
}
