/*
 * @Author: JimZhang
 * @Date: 2025-07-18 10:54:21
 * @LastEditors: 很拉风的James
 * @LastEditTime: 2025-07-21 23:44:06
 * @FilePath: /server/api/user_api.go
 * @Description:
 *
 */
package api

import (
	"fmt"
	"server/service"
	"server/service/dto"
	"server/utils"

	"github.com/gin-gonic/gin"
)

const (
	ERR_CODE_ADD_USER       = 10011
	ERR_CODE_GET_USER_BY_ID = 10012
	ERR_CODE_GET_USER_LIST  = 10013
	ERR_CODE_UPDATE_USER    = 10014
	ERR_CODE_DELETE_USER    = 10015
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

	iUser, err := m.Service.Login(iUserLoginDTO)

	if err != nil {
		m.Fail(ResponseJson{
			Msg: err.Error(),
		})
		return
	}

	token, _ := utils.GenerateToken(int(iUser.Id), iUser.Username)

	m.OK(ResponseJson{
		Msg: "Login Success",
		Data: gin.H{
			"token": token,
			"user":  iUser,
		},
	})
	// fmt.Println("login is called")
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"msg": "login Success",
	// })

	// OK(ctx, ResponseJson{
	// 	Msg: "login Success",
	// })

	// Fail(ctx, ResponseJson{
	// 	Code: 9001,
	// 	Msg:  "login Fail",
	// })
}

func (m UserApi) AddUser(c *gin.Context) {
	var iUserAddDTO dto.UserAddDTO
	if err := m.BulidRequest(BulidRequestOptions{
		Ctx: c,
		DTO: &iUserAddDTO,
	}).GetError(); err != nil {
		return
	}

	file, _ := c.FormFile("file")
	stFilePath := fmt.Sprintf("./uploads/%s", file.Filename)
	_ = c.SaveUploadedFile(file, stFilePath)
	iUserAddDTO.Avatar = stFilePath

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
