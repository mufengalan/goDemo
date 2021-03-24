package model

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"head_server/support"
)

var connection = "db_ex_doctor"

func SetChainBranchRelation() *xorm.Engine {

	host := support.GetCfgString("LOCAL_HOST")
	port := support.GetCfgInt("LOCAL_PORT")
	username := support.GetCfgString("LOCAL_USER")
	password := support.GetCfgString("LOCAL_PASS")
	dbname := connection
	addr := fmt.Sprintf("%s:%d", host, port)
	return support.OpenEngine(addr, username, password, dbname)
}

type student struct {
	Name string
}

func GetShopList(AppId string) string {

	engine := SetChainBranchRelation()
	engine.ShowSQL(true)
	Student := &student{}
	has, _ := engine.Where("id = ?", 1).Cols("name").Get(Student)
	//has, _:= engine.Sql("select head_app_id from t_chain_branch_relation where id = ?",2).Get(student)
	fmt.Printf("user:%s\n", Student.Name)

	if has {
		fmt.Println("查询成功")

	} else {
		fmt.Println("数据不存在")
	}
	return Student.Name
}
