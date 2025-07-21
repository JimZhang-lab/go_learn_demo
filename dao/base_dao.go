/*
 * @Author: JimZhang
 * @Date: 2025-07-20 12:07:22
 * @LastEditors: 很拉风的James
 * @LastEditTime: 2025-07-20 12:15:13
 * @FilePath: /server/dao/base_dao.go
 * @Description:
 *
 */
package dao

import (
	"server/global"

	"gorm.io/gorm"
)

type BaseDao struct {
	Orm *gorm.DB
}

func NewBaseDao() BaseDao {
	return BaseDao{
		Orm: global.DB,
	}
}
