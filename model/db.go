package model

import (
	"admin/utils"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

//打开数据库
func InitDb() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("数据库信息错误", err)
		return db, err
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("数据库连接错误", err)
		return db, err
	}
	return db, err
}
