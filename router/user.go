/*
 * @Author: JimZhang
 * @Date: 2025-06-30 23:43:19
 * @LastEditors: 很拉风的James
 * @LastEditTime: 2025-07-21 23:36:37
 * @FilePath: /server/router/user.go
 * @Description:
 *
 */
package router

import (
	"server/api"

	"github.com/gin-gonic/gin"
)

func InitUserRouters() {
	RegisterRouter(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		userApi := api.New_UserApi()
		rgPublicUser := rgPublic.Group("user").Use(func() gin.HandlerFunc {
			return func(ctx *gin.Context) {
				// ctx.AbortWithStatusJSON(200, gin.H{
				// 	"msg": "login Middleware",
				// })
			}
		}())
		{
			rgPublicUser.POST("/login", userApi.Login)
		}

		rgAuthUser := rgAuth.Group("user")
		{
			rgAuthUser.POST("", userApi.AddUser)
			rgAuthUser.GET("/:id", userApi.SearchUserById)
			rgAuthUser.POST("/list", userApi.GetUserList)
			rgAuthUser.PUT("/:id", userApi.UpdateUser)
			rgAuthUser.DELETE("/:id", userApi.DeleteUserById)
		}
	})
}
