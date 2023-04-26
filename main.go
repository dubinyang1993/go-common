package main

import (
	"fmt"
	"go-common/config"
	"go-common/mysql"
	"go-common/redis"
)

func main() {
	fmt.Println("go-common is starting...")

	//配置文件和日志
	config.Init()

	//初始化数据库
	mysql.Init()

	//初始化Redi
	redis.Init()
}
