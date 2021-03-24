package model

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"head_server/support"
)

var connection = "db_ex_chain_store"

func SetChainBranchRelation() *xorm.Engine {
	host := "127.0.0.1"
	port := "3306"
	username := "root"
	password := "123456"
	dbname := connection
	fmt.Printf("====%v \n", host)
	addr := fmt.Sprintf("%s:%d", host, port)
	return support.OpenEngine(addr, username, password, dbname)
}

type student struct {
	name string
}

func GetShopList(AppId string) string {

	engine := SetChainBranchRelation()
	engine.ShowSQL(true)
	Student := &student{}
	has, _ := engine.Where("id = ?", 1).Cols("name").Get(Student)
	//has, _:= engine.Sql("select head_app_id from t_chain_branch_relation where id = ?",2).Get(student)
	fmt.Printf("user:%v\n", Student.name)

	if has {
		return Student.name
	} else {
		fmt.Println("数据不存在")
		return Student.name
	}
}
