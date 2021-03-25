package model

import (
	"fmt"
	"head_server/database"
)

var doctor = database.Connection{
	Host:     "LOCAL_HOST",
	Port:     "LOCAL_PORT",
	Username: "LOCAL_USER",
	Password: "LOCAL_PASS",
	Dbname:   "db_ex_doctor",
}

type student struct {
	Name string
}

func GetDoctorList(AppId string) string {
	engine := MysqlConnection(doctor)
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
