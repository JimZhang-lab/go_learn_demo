/*
 * @Author: JimZhang
 * @Date: 2025-07-19 16:37:10
 * @LastEditors: 很拉风的James
 * @LastEditTime: 2025-07-22 12:16:20
 * @FilePath: /server/service/dto/user_dto.go
 * @Description:
 *
 */
package dto

import (
	"server/model"
)

type UserLoginDTO struct {
	Username string `json:"username" binding:"required" message:"用户名不能为空"`
	Password string `json:"password" binding:"required" message:"密码不能为空"`
}

// 添加用户 DTO
type UserAddDTO struct {
	Id       int64
	Username string `json:"username" form:"username" binding:"required" message:"用户名不能为空"`
	Password string `json:"password,omitempty" form:"password" binding:"required" message:"密码不能为空"`
	Realname string `json:"realname" form:"realname"`
	Avatar   string
	Mobile   string `json:"mobile" form:"mobile"`
	Email    string `json:"email" form:"email"`
}

func (m *UserAddDTO) ConvertToModel(iUser *model.User) {
	iUser.Username = m.Username
	// stHash, _ := utils.Encrpypt(m.Password)
	// iUser.Password = stHash
	iUser.Password = m.Password
	iUser.Realname = m.Realname
	iUser.Avatar = m.Avatar
	iUser.Mobile = m.Mobile
	iUser.Email = m.Email
}

// 用户列表相关 DTO
type UserListDTO struct {
	PaginateDTO
}

// 用户更新 DTO
type UserUpdateDTO struct {
	Id       int64  `json:"id" form:"id" uri:"id"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Realname string `json:"realname" form:"realname"`
	Avatar   string `json:"avatar" form:"avatar"`
	Mobile   string `json:"mobile" form:"mobile"`
	Email    string `json:"email" form:"email"`
}

func (m *UserUpdateDTO) ConvertToModel(iUser *model.User) {
	iUser.Username = m.Username
	// stHash, _ := utils.Encrpypt(m.Password)
	// iUser.Password = stHash
	iUser.Password = m.Password
	iUser.Realname = m.Realname
	iUser.Avatar = m.Avatar
	iUser.Mobile = m.Mobile
	iUser.Email = m.Email
}
