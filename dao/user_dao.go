/*
 * @Author: JimZhang
 * @Date: 2025-07-20 12:07:52
 * @LastEditors: 很拉风的James
 * @LastEditTime: 2025-07-21 23:30:07
 * @FilePath: /server/dao/user_dao.go
 * @Description:
 *
 */
package dao

import (
	"server/model"
	"server/service/dto"
)

var userDao *UserDao

type UserDao struct {
	BaseDao
}

func NewUserDao() *UserDao {
	if userDao == nil {
		userDao = &UserDao{
			NewBaseDao(),
		}
	}
	return userDao
}

func (m *UserDao) GetUserByNameAndPassword(stUserName, stPassword string) model.User {
	var iUser model.User
	m.Orm.Model(&iUser).Where("username=? and password=?", stUserName, stPassword).Find(&iUser)
	return iUser
}

func (m *UserDao) CheckUserNameExist(stUserName string) bool {
	var nTotal int64
	m.Orm.Model(&model.User{}).Where("username=?", stUserName).Count(&nTotal)
	return nTotal > 0
}

func (m *UserDao) AddUser(iUserAddDTO *dto.UserAddDTO) error {
	var iUser model.User
	iUserAddDTO.ConvertToModel(&iUser)

	err := m.Orm.Save(&iUser).Error
	if err == nil {
		iUserAddDTO.Id = iUser.Id
		iUserAddDTO.Password = ""
	}
	return err
}

func (m *UserDao) GetUserById(nId int64) (model.User, error) {
	var iUser model.User
	err := m.Orm.Where("id=?", nId).First(&iUser).Error
	return iUser, err
}

func (m *UserDao) GetUserList(iUserListDTO *dto.UserListDTO) ([]model.User, int64, error) {
	var giUserList []model.User // 当页户列表
	var nTotal int64            // 总数

	err := m.Orm.Model(&model.User{}).Scopes(Paginate(iUserListDTO.PaginateDTO)).
		Find(&giUserList).
		Offset(-1).Limit(-1).
		Count(&nTotal).Error

	return giUserList, nTotal, err
}

func (m *UserDao) UpdateUser(iUserUpdateDTO *dto.UserUpdateDTO) error {
	var iUser model.User

	// save 默认清空之前的数据，所以要先查出来赋值后再修改
	// m.Orm.Where("id=?", iUserUpdateDTO.Id).First(&iUser)
	m.Orm.First(&iUser, iUserUpdateDTO.Id)
	// fmt.Println(iUser)
	iUserUpdateDTO.ConvertToModel(&iUser)

	return m.Orm.Save(&iUser).Error
}

func (m *UserDao) DeleteUserById(nId int64) error {
	return m.Orm.Delete(&model.User{}, nId).Error
}
