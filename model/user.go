/*
 * @Author: JimZhang
 * @Date: 2025-07-18 13:53:58
 * @LastEditors: 很拉风的James
 * @LastEditTime: 2025-07-22 13:33:28
 * @FilePath: /server/model/user.go
 * @Description:
 *
 */
package model

import (
	"server/utils"

	"gorm.io/gorm"
)

type User struct {
	Id       int64  `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Username string `gorm:"column:username;type:varchar(24);not null" json:"username"`
	Password string `gorm:"column:password;type:varchar(96);not null" json:"-"`
	Realname string `gorm:"column:realname;type:varchar(255)" json:"realname"`
	Avatar   string `gorm:"column:avatar;type:varchar(255)" json:"avatar"`
	Mobile   string `gorm:"column:mobile;type:varchar(11)" json:"mobile"`
	Email    string `gorm:"column:email;type:varchar(128)" json:"email"`
}

// Encrpypt 加密用户密码
// 函数将用户模型中的密码进行加密处理
// 返回加密过程中可能出现的错误
func (m *User) Encrpypt() error {
	stHash, err := utils.Encrpypt(m.Password)

	if err == nil {
		m.Password = stHash
	}
	return err
}

// BeforeCreate 在创建用户前的钩子函数
// 用于在创建用户记录前自动加密密码
// 参数:
//
//	orm *gorm.DB - GORM 数据库实例
//
// 返回创建前处理过程中可能出现的错误
func (m *User) BeforeCreate(orm *gorm.DB) error {
	return m.Encrpypt()
}

// 用户登录信息
type LoginUser struct {
	Id       int64
	Username string
}
