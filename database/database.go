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

//本地库测试
var Doctor = Connection{
	Host:     "LOCAL_HOST",
	Port:     "LOCAL_PORT",
	Username: "LOCAL_USER",
	Password: "LOCAL_PASS",
	Dbname:   "db_ex_doctor",
}
