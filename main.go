package main

import (
	"admin/model"
	"admin/routes"
	"fmt"
)

func main() {

	//初始化数据库
	DB, err := model.InitDb()
	if err != nil {
		fmt.Println("数据库初始化失败", err)
		return
	}

	//延迟数据库关闭
	defer DB.Close()

	//初始化路由
	routes.InitRouter()
}
