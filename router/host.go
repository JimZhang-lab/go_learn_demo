/*
 * @Author: JimZhang
 * @Date: 2025-06-30 23:43:19
 * @LastEditors: 很拉风的James
 * @LastEditTime: 2025-07-20 17:52:04
 * @FilePath: /server/router/host.go
 * @Description:
 *
 */
package router

import (
	"server/api"

	"github.com/gin-gonic/gin"
)

func InitHostRouters() {
	RegisterRouter(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		hostApi := api.NewHostApi()

		rgAuthHost := rgAuth.Group("host")
		{
			rgAuthHost.POST("/shutdown", hostApi.Shutdown)
		}
	})
}
