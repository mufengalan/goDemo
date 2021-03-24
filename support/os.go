package support

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"os"
	"strconv"
)

/**读取env中的值 返回string类型*/
func GetCfgString(key string, def ...string) string {
	var defaultV string
	if len(def) <= 0 {
		defaultV = ""
	} else {
		defaultV = def[0]
	}
	value := os.Getenv(key)
	if value == "" {
		value = defaultV
	}
	return value
}

/**读取env中的值 返回int类型*/
func GetCfgInt(key string, def ...int) int {
	var defaultV int
	if len(def) <= 0 {
		defaultV = 0
	} else {
		defaultV = def[0]
	}
	value := os.Getenv(key)
	v, err := strconv.Atoi(value)
	if err != nil {
		v = defaultV
	}
	return v
}

/**建立数据库链接*/
func OpenEngine(addr, username, password, dbname string) *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, addr, dbname))
	if err != nil {
		panic(err)
	}
	if err := engine.Ping(); err != nil {
		fmt.Println("数据库链接失败")
		fmt.Println(err)
	}
	//defer engine.Close() //延迟关闭数据库
	fmt.Println("数据库链接成功")
	return engine
}
