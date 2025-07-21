/*
 * @Author: JimZhang
 * @Date: 2025-06-30 23:04:48
 * @LastEditors: 很拉风的James
 * @LastEditTime: 2025-07-20 10:59:35
 * @FilePath: /server/cmd/cmd.go
 * @Description:
 *
 */
package cmd

import (
	"fmt"
	"server/conf"
	"server/global"
	"server/router"
	"server/utils"
)

func Start() {
	var initErr error

	// 初始化系统配置
	conf.InitConfig()

	// 初始化日志
	global.Logger = conf.InitLogger()

	// 初始化数据库连接
	db, err := conf.InitMysql()
	global.DB = db

	if err != nil {
		initErr = utils.AppendError(initErr, err)
	}

	// 初始化 redis 连接
	rdClient, err := conf.InitRedis()
	global.RedisClient = rdClient

	if err != nil {
		initErr = utils.AppendError(initErr, err)
	}
	// _ = global.RedisClient.Delete("test", "addr")
	// fmt.Println(global.RedisClient.Get("test"))

	// 初始化过程中存在的错误
	if initErr != nil {
		if global.Logger != nil {
			global.Logger.Error(initErr.Error())
		}
		panic(initErr.Error())
	}
	// 初始化路由
	router.InitRouter()
}

func Clean() {

	fmt.Println("================= Clean =================")
}
