package model

import (
	"fmt"
	"head_server/database"
)

type Student struct {
	Name string
}

func GetDoctorList(id int) string {
	//连接数据库
	engine := MysqlConnection(database.Doctor)
	engine.ShowSQL(true)
	student := &Student{}
	fmt.Printf("==%d", id)
	has, _ := engine.Where("id = ?", id).Cols("name").Get(student)
	//has, _:= engine.Sql("select head_app_id from t_chain_branch_relation where id = ?",2).Get(student)
	fmt.Printf("user:%s\n", student.Name)
	if has {
		fmt.Println("查询成功")

	} else {
		fmt.Println("数据不存在")
	}
	return student.Name
}
