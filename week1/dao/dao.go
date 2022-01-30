package dao

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var ErrDataNotFound = errors.New("dao: data not found")

type sqlDAO struct {
	db *sql.DB
}

func NewMySQLDao(host string, port int, user, passwd, dbname string) (daoInstance *sqlDAO, err error) {
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", user, passwd, host, port, dbname)
	db, err := sql.Open("mysql", dbDSN)
	if err != nil {
		return
	}

	daoInstance = &sqlDAO{db}
	return
}
