/*
 * @Author: JimZhang
 * @Date: 2025-07-20 17:00:56
 * @LastEditors: 很拉风的James
 * @LastEditTime: 2025-07-20 17:09:20
 * @FilePath: /server/middleware/cors.go
 * @Description:
 *
 */
package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	cfg := cors.Config{
		// 允许的请求来源
		// AllowOrigins: []string{"https://foo.com"},
		// 允许的请求方法
		AllowMethods: []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS", "HEAD"},
		// 请求头
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Accept", "token"},
		// 请求头，哪些允许被客户端读取
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, // 允许带cookie的请求
		// 允许的请求来源
		AllowOriginFunc: func(origin string) bool {
			// 设置为 true 表示任何站点都可以访问
			return true
			// return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}

	return cors.New(cfg)
}
