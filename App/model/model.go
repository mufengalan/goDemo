package model

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"head_server/support"
)

type Connection struct {
	host     string
	port     int
	username string
	password string
	dbname   string
}

func MysqlConnection(conn Connection) *xorm.Engine {
	host := support.GetCfgString("LOCAL_HOST")
	port := support.GetCfgInt("LOCAL_PORT")
	username := support.GetCfgString("LOCAL_USER")
	password := support.GetCfgString("LOCAL_PASS")
	dbname := connection
	addr := fmt.Sprintf("%s:%d", host, port)
	return support.OpenEngine(addr, username, password, dbname)
}
