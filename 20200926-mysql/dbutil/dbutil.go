package dbutil

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	err := initDB()
	if err != nil {
		fmt.Printf("连接数据库失败！！！,err: %v", err)
	}
}

// 定义一个全局对象db
var DB *sql.DB

// 定义一个初始化数据库的函数
func initDB() (err error) {
	// DSN:Data Source Name
	dsn := "root:root1234@tcp(192.168.0.27:3306)/sql_test?charset=utf8mb4&parseTime=True"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = DB.Ping()
	if err != nil {
		return err
	}
	return nil
}
