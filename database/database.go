package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"head_server/app/model"
)

type Database struct {
	Example *xorm.Engine
}

var DB *Database

func Load() {
	DB = &Database{
		Example: model.SetChainBranchRelation(),
	}
}
