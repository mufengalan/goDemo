package model

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"head_server/database"
	"head_server/support"
)

func MysqlConnection(conn database.Connection) *xorm.Engine {
	host := support.GetCfgString(conn.Host)
	port := support.GetCfgInt(conn.Port)
	username := support.GetCfgString(conn.Username)
	password := support.GetCfgString(conn.Password)
	dbname := conn.Dbname
	addr := fmt.Sprintf("%s:%d", host, port)
	return support.OpenEngine(addr, username, password, dbname)
}
