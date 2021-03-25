package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var DB *Database

type Database struct {
	Example *xorm.Engine
}

type Connection struct {
	Host     string
	Port     string
	Username string
	Password string
	Dbname   string
}
