/*
 * @Author: JimZhang
 * @Date: 2025-07-20 12:06:06
 * @LastEditors: 很拉风的James
 * @LastEditTime: 2025-07-21 23:30:32
 * @FilePath: /server/service/user_service.go
 * @Description:
 *
 */
package service

import (
	"errors"
	"server/dao"
	"server/model"
	"server/service/dto"
)

var userService *UserService

type UserService struct {
	BaseService
	Dao *dao.UserDao
}

func NewUserService() *UserService {
	if userService == nil {
		userService = &UserService{
			Dao: dao.NewUserDao(),
		}
	}
	return userService
}

func (m *UserService) Login(iUserDTO dto.UserLoginDTO) (model.User, error) {
	var errResult error

	iUser := m.Dao.GetUserByNameAndPassword(iUserDTO.Username, iUserDTO.Password)
	if iUser.Id == 0 {
		errResult = errors.New("用户名或密码错误")
	}

	return iUser, errResult
}

func (m *UserService) AddUser(iUserAddDTO *dto.UserAddDTO) error {
	if m.Dao.CheckUserNameExist(iUserAddDTO.Username) {
		return errors.New("用户名已存在")
	}
	return m.Dao.AddUser(iUserAddDTO)
}

func (m *UserService) GetUserById(iCommonIDDTO *dto.CommonIDDTO) (model.User, error) {
	return m.Dao.GetUserById(iCommonIDDTO.Id)
}

func (m *UserService) GetUserList(iUserListDTO *dto.UserListDTO) ([]model.User, int64, error) {
	return m.Dao.GetUserList(iUserListDTO)
}

func (m *UserService) UpdateUser(iUserUpdateDTO *dto.UserUpdateDTO) error {
	if iUserUpdateDTO.Id <= 0 {
		return errors.New("用户ID不能为空")
	}

	return m.Dao.UpdateUser(iUserUpdateDTO)
}

func (m *UserService) DeleteUserById(iCommonIDDTO *dto.CommonIDDTO) error {
	if iCommonIDDTO.Id <= 0 {
		return errors.New("用户ID不能为空")
	}

	return m.Dao.DeleteUserById(iCommonIDDTO.Id)
}
