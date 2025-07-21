/*
 * @Author: JimZhang
 * @Date: 2025-07-18 13:53:58
 * @LastEditors: 很拉风的James
 * @LastEditTime: 2025-07-21 23:26:49
 * @FilePath: /server/model/user.go
 * @Description:
 *
 */
package model

type User struct {
	Id       int64  `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Username string `gorm:"column:username;type:varchar(24);not null" json:"username"`
	Password string `gorm:"column:password;type:varchar(48);not null" json:"-"`
	Realname string `gorm:"column:realname;type:varchar(255)" json:"realname"`
	Avatar   string `gorm:"column:avatar;type:varchar(255)" json:"avatar"`
	Mobile   string `gorm:"column:mobile;type:varchar(11)" json:"mobile"`
	Email    string `gorm:"column:email;type:varchar(128)" json:"email"`
}
