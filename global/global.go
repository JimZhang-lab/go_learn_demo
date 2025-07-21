/*
 * @Author: JimZhang
 * @Date: 2025-07-18 11:51:24
 * @LastEditors: 很拉风的James
 * @LastEditTime: 2025-07-19 12:26:40
 * @FilePath: /server/global/global.go
 * @Description:
 *
 */
package global

import (
	"server/conf"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// type RedisClientInterface interface {
// 	Set(key string, value any) error
// 	Get(key string) (string, error)
// 	Delete(key ...string) error
// }

var (
	Logger      *zap.SugaredLogger
	DB          *gorm.DB
	RedisClient *conf.RedisClient
)
