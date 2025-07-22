/*
 * @Author: JimZhang
 * @Date: 2025-07-20 12:06:06
 * @LastEditors: 很拉风的James
 * @LastEditTime: 2025-07-22 13:26:49
 * @FilePath: /server/service/user_service.go
 * @Description:
 *
 */
package service

import (
	"errors"
	"fmt"
	"server/dao"
	"server/global"
	"server/global/constants"
	"server/model"
	"server/service/dto"
	"server/utils"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"
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

func GenerateAndCacheUserTokenToRedis(nUserId int64, stUserName string) (string, error) {
	token, err := utils.GenerateToken(int(nUserId), stUserName)
	if err == nil {
		err = global.RedisClient.Set(strings.Replace(constants.LOGIN_USER_TOKEN_REDIS_KEY, "{ID}", strconv.Itoa(int(nUserId)), -1), token, viper.GetDuration("jwt.tokenExpire")*time.Minute)
	}
	return token, err
}

func (m *UserService) Login(iUserDTO dto.UserLoginDTO) (model.User, string, error) {
	var errResult error
	var token string

	iUser, err := m.Dao.GetUserByName(iUserDTO.Username)

	// 用户名或密码不正确
	if err != nil || !utils.CompareHashAndPassword(iUser.Password, iUserDTO.Password) {
		errResult = errors.New("用户名或密码错误")
	} else {
		// 登录成功，token 生成
		token, err = GenerateAndCacheUserTokenToRedis(iUser.Id, iUser.Username)
		if err != nil {
			errResult = errors.New(fmt.Sprintf("Generate Token Error: %s", err.Error()))
		}
	}

	return iUser, token, errResult
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
