/*
 * @Author: JimZhang
 * @Date: 2025-06-29 01:21:25
 * @LastEditors: 很拉风的James
 * @LastEditTime: 2025-07-22 13:01:37
 * @FilePath: /server/router/router.go
 * @Description:
 *
 */
package router

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"strings"
	"syscall"
	"time"

	_ "server/docs"
	"server/global"
	"server/middleware"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type IFnRegisterRouter = func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup)

var (
	gfnRoutes []IFnRegisterRouter
)

func RegisterRouter(fn IFnRegisterRouter) {
	if fn == nil {
		return
	}
	gfnRoutes = append(gfnRoutes, fn)
}

func InitRouter() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	r := gin.Default()
	r.Use(middleware.Cors())
	rgPublic := r.Group("/api/v1/public")
	rgAuth := r.Group("/api/v1")
	rgAuth.Use(middleware.Auth())

	initBasePlatformRouters()

	// 注册自定义校验器
	registerCustValidator()

	for _, fnRegisterRouter := range gfnRoutes {
		fnRegisterRouter(rgPublic, rgAuth)
	}

	// 集成 swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	stPort := viper.GetString("server.port")
	if stPort == "" {
		stPort = "8888"
	}

	//
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", stPort),
		Handler: r,
	}

	// 协程启动
	go func() {
		global.Logger.Info(fmt.Sprintf("Start Server Listen: %s", stPort))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// log.Fatalf("listen: %s\n", err)
			// TODO: 记录日志
			global.Logger.Error(fmt.Sprintf("Start server error: %s", err.Error()))
			// fmt.Println(fmt.Sprintf("Start server error: %s", err.Error()))
			return
		}

		// fmt.Println(fmt.Sprintf("Start Server Listen: %s", stPort))
	}()

	<-ctx.Done()

	ctx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()
	if err := srv.Shutdown(ctx); err != nil {
		// log.Fatalf("Server Shutdown: %s", err)
		// TODO: 错误日志
		global.Logger.Error(fmt.Sprintf("Server Shutdown: %s", err.Error()))
		// fmt.Println(fmt.Sprintf("Server Shutdown: %s", err))
		return
	}
	global.Logger.Info("Server exiting")
	// fmt.Println("Server exiting")
}

func initBasePlatformRouters() {
	InitUserRouters()
	InitHostRouters()
}

// 注册自定义验证器
func registerCustValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("first_is_a", func(fl validator.FieldLevel) bool {
			if value, ok := fl.Field().Interface().(string); ok {
				if value != "" && (strings.Index(value, "a") == 0) {
					return true
				}
			}
			return false
		})
	}
}
