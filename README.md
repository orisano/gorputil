# gorputil
[![Build Status](https://travis-ci.com/orisano/gorputil.svg?branch=master)](https://travis-ci.com/orisano/gorputil)
[![Maintainability](https://api.codeclimate.com/v1/badges/7b6fd84c34e72fdd81d4/maintainability)](https://codeclimate.com/github/orisano/gorputil/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/7b6fd84c34e72fdd81d4/test_coverage)](https://codeclimate.com/github/orisano/gorputil/test_coverage)

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
	slaves := []*gorp.DbMap{
		{Db: slaveDb1, Dialect: gorp.MySQLDialect{}},
		{Db: slaveDb2, Dialect: gorp.MySQLDialect{}},
	}
	db := gorputil.NewClusterMap(master, slaves, &gorputil.Sequential{})

	db.Exec("insert into users(name, age, weight) values(foo, 25, 58.0)") // master
	db.Query("select * from users")                                       // slave1
	db.Query("select * from users")                                       // slave2
	db.Query("select * from users")                                       // slave1
	db.Query("delete from users")                                         // master
}
```

## Author
Nao Yonashiro (@orisano)

## LICENSE
MIT
