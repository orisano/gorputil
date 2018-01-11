# gorputil
`gopkg.in/gorp.v2` utility.

## Installation
`go get -u github.com/orisano/gorputil`

## How to Use
```go
package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/orisano/gorputil"
	"gopkg.in/gorp.v2"
)

func main() {
	masterDb, _ := sql.Open("mysql", "user:password@tcp(master.mysql.example)/dbname")
	slaveDb1, _ := sql.Open("mysql", "user:password@tcp(slave1.mysql.example)/dbname")
	slaveDb2, _ := sql.Open("mysql", "user:password@tcp(slave2.mysql.example)/dbname")

	master := &gorp.DbMap{Db: masterDb, Dialect: gorp.MySQLDialect{}}
	slaves := []gorp.SqlExecutor{
		&gorp.DbMap{Db: slaveDb1, Dialect: gorp.MySQLDialect{}},
		&gorp.DbMap{Db: slaveDb2, Dialect: gorp.MySQLDialect{}},
	}
	db := gorputil.BalancedSqlExecutor(master, slaves, &gorputil.Sequential{})

	db.Exec("insert into users(name, age, weight) values(foo, 25, 58.0)") // master
	db.Query("select * from users")                                       // slave1
	db.Query("select * from users")                                       // slave2
	db.Query("select * from users")                                       // slave1
	db.Query("delete from users")                                         // master
}
```

## LICENSE
MIT