/*
 * @Author: JimZhang
 * @Date: 2025-07-18 10:54:21
 * @LastEditors: 很拉风的James
 * @LastEditTime: 2025-07-22 13:35:45
 * @FilePath: /server/api/user_api.go
 * @Description:
 *
 */
package api

import (
	"fmt"
	"net/http"
	"server/service"
	"server/service/dto"

	"github.com/gin-gonic/gin"
)

const (
	ERR_CODE_ADD_USER       = 10011
	ERR_CODE_GET_USER_BY_ID = 10012
	ERR_CODE_GET_USER_LIST  = 10013
	ERR_CODE_UPDATE_USER    = 10014
	ERR_CODE_DELETE_USER    = 10015
	ERR_CODE_LOGIN_USER     = 10016
)

type UserApi struct {
	BaseApi
	Service *service.UserService
}

func New_UserApi() UserApi {
	return UserApi{
		BaseApi: NewBaseApi(),
		Service: service.NewUserService(),
	}
}

// @Tags 用户管理
// @Summary 登录
// @Description 登录
// @Param name formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {object} string "{"msg":"登录成功"}"
// @Failure 400 {object} string "{"msg":"登录失败"}"
// @Router /api/v1/public/user/login [post]
func (m UserApi) Login(c *gin.Context) {

	var iUserLoginDTO dto.UserLoginDTO

	if err := m.BulidRequest(BulidRequestOptions{Ctx: c, DTO: &iUserLoginDTO}).GetError(); err != nil {

		return
	}

	iUser, token, err := m.Service.Login(iUserLoginDTO)

	if err != nil {
		m.Fail(ResponseJson{
			Status: http.StatusUnauthorized,
			Code:   ERR_CODE_LOGIN_USER,
			Msg:    err.Error(),
		})
		return
	}

	// 将用户ID和token存储到Redis中
	m.OK(ResponseJson{
		Msg: "Login Success",
		Data: gin.H{
			"token": token,
			"user":  iUser,
		},
	})

}

// @Tags 用户管理
// @Summary 添加用户
// @Description 创建新用户
// @Param user body dto.UserAddDTO true "用户数据"
// @Success 200 {object} ResponseJson "用户添加成功"
// @Failure 400 {object} ResponseJson "用户添加失败"
// @Router /api/v1/user [post]
func (m UserApi) AddUser(c *gin.Context) {
	var iUserAddDTO dto.UserAddDTO
	if err := m.BulidRequest(BulidRequestOptions{
		Ctx: c,
		DTO: &iUserAddDTO,
	}).GetError(); err != nil {
		return
	}

	// file, _ := c.FormFile("file")
	// stFilePath := fmt.Sprintf("./uploads/%s", file.Filename)
	// _ = c.SaveUploadedFile(file, stFilePath)
	// iUserAddDTO.Avatar = stFilePath

	err := m.Service.AddUser(&iUserAddDTO)

	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_ADD_USER,
			Msg:  err.Error(),
		})
		return
	}
	m.OK(ResponseJson{
		Data: iUserAddDTO,
	})
}

// @Tags 用户管理
// @Summary 根据ID获取用户
// @Description 获取指定ID的用户信息
// @Param id path int true "用户ID"
// @Success 200 {object} ResponseJson "用户信息"
// @Failure 400 {object} ResponseJson "获取用户信息失败"
// @Router /api/v1/user/{id} [get]
func (m UserApi) SearchUserById(c *gin.Context) {
	var iCommonIDDTO dto.CommonIDDTO
	if err := m.BulidRequest(BulidRequestOptions{
		Ctx:     c,
		DTO:     &iCommonIDDTO,
		BindUri: true,
	}).GetError(); err != nil {
		return
	}
	user, err := m.Service.GetUserById(&iCommonIDDTO)
	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_GET_USER_BY_ID,
			Msg:  err.Error(),
		})
		return
	}
	m.OK(ResponseJson{
		Data: user,
	})
}

// @Tags 用户管理
// @Summary 获取用户列表
// @Description 分页获取用户列表
// @Param page query int false "页码"
// @Param size query int false "每页数量"
// @Success 200 {object} ResponseJson "用户列表"
// @Failure 400 {object} ResponseJson "获取用户列表失败"
// @Router /api/v1/user [get]
func (m UserApi) GetUserList(c *gin.Context) {
	var iUserListDTO dto.UserListDTO
	if err := m.BulidRequest(BulidRequestOptions{
		Ctx:     c,
		DTO:     &iUserListDTO,
		BindUri: true,
	}).GetError(); err != nil {
		return
	}
	userList, total, err := m.Service.GetUserList(&iUserListDTO)
	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_GET_USER_LIST,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Data:  userList,
		Total: total,
	})
}

// @Tags 用户管理
// @Summary 更新用户
// @Description 更新指定ID的用户信息
// @Param id path int true "用户ID"
// @Param user body dto.UserUpdateDTO true "用户数据"
// @Success 200 {object} ResponseJson "用户更新成功"
// @Failure 400 {object} ResponseJson "用户更新失败"
// @Router /api/v1/user/{id} [put]
func (m UserApi) UpdateUser(c *gin.Context) {
	var iUserUpdateDTO dto.UserUpdateDTO
	// 这里得到的 id 可能是 string 类型，需要转换成 int64
	// strId := c.Param("id")
	// iUserUpdateDTO.Id, _ = strconv.ParseInt(strId, 10, 64)
	if err := m.BulidRequest(BulidRequestOptions{
		Ctx:     c,
		DTO:     &iUserUpdateDTO,
		BindAll: true, // 确保从 URI 绑定参数
	}).GetError(); err != nil {
		return
	}

	fmt.Println(iUserUpdateDTO)
	if err := m.Service.UpdateUser(&iUserUpdateDTO); err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_UPDATE_USER,
			Msg:  err.Error(),
		})
		return
	}
	m.OK(ResponseJson{
		Code: ERR_CODE_UPDATE_USER,
		Msg:  "更新用户成功",
	})
}

// @Tags 用户管理
// @Summary 根据ID删除用户
// @Description 删除指定ID的用户
// @Param id path int true "用户ID"
// @Success 200 {object} ResponseJson "用户删除成功"
// @Failure 400 {object} ResponseJson "用户删除失败"
// @Router /api/v1/user/{id} [delete]
func (m *UserApi) DeleteUserById(c *gin.Context) {
	var iCommonIDDTO dto.CommonIDDTO
	if err := m.BulidRequest(BulidRequestOptions{
		Ctx:     c,
		DTO:     &iCommonIDDTO,
		BindUri: true, // 确保从 URI 绑定参数
	}).GetError(); err != nil {
		return
	}
	user, err := m.Service.GetUserById(&iCommonIDDTO)
	if err != nil {
		m.Fail(ResponseJson{
			Code: ERR_CODE_DELETE_USER,
			Msg:  err.Error(),
		})
		return
	}

	if err := m.Service.DeleteUserById(&iCommonIDDTO); err != nil {
		m.Fail(ResponseJson{
			Code: ERR_CODE_DELETE_USER,
			Msg:  err.Error(),
		})
		return
	}
	m.OK(ResponseJson{
		Msg:  "删除成功",
		Data: user,
	})

}
