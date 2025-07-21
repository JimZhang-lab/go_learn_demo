/*
 * @Author: JimZhang
 * @Date: 2025-06-30 23:04:58
 * @LastEditors: 很拉风的James
 * @LastEditTime: 2025-07-19 12:58:40
 * @FilePath: /server/main.go
 * @Description:
 *
 */
package main

import (
	"server/cmd"
)

// @title Go admin 开发
// @version 1.0
// @description Go admin 测试
func main() {

	defer cmd.Clean()

	cmd.Start()

}
