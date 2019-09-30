package main

import (
	"fmt"
	"github.com/lexkong/log"
	"go-common/config"
	"go-common/db"
	"go-common/redis"
)

func main() {
	//配置文件和日志
	config.Init()
	fmt.Println("go-common is starting...")
	log.Info("go-common is starting...")

	//初始化数据库
	db.Init()

	//初始化Redi
	redis.Init()
}
