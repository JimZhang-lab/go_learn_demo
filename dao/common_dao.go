/*
 * @Author: JimZhang
 * @Date: 2025-07-21 20:14:42
 * @LastEditors: 很拉风的James
 * @LastEditTime: 2025-07-21 20:18:58
 * @FilePath: /server/dao/common_dao.go
 * @Description:
 *
 */
package dao

import (
	"server/service/dto"

	"gorm.io/gorm"
)

// Paginate 是一个分页函数，用于在数据库查询中应用分页逻辑。
// 参数:
//   p - PaginateDTO 类型，包含分页信息（如页码和每页记录数）。
// 返回值:
//   一个函数，该函数接受 *gorm.DB 类型的数据库连接，并返回应用了分页逻辑的 *gorm.DB 类型结果。
func Paginate(p dto.PaginateDTO) func(orm *gorm.DB) *gorm.DB {
	return func(orm *gorm.DB) *gorm.DB {
		return orm.Offset((p.GetPage() - 1) * p.GetLimit()).Limit(p.GetLimit())
	}
}
