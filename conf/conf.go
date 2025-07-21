/*
 * @Author: JimZhang
 * @Date: 2025-06-30 23:09:14
 * @LastEditors: 很拉风的James
 * @LastEditTime: 2025-06-30 23:18:59
 * @FilePath: /vue_learn/server/conf/conf.go
 * @Description:
 *
 */
package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf/")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Fatal error config file: %s \n", err.Error()))
	}
	fmt.Println(viper.GetString("server.port"))
}
