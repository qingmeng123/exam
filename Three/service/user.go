/*******
* @Author:qingmeng
* @Description:
* @File:user
* @Date2021/12/10
 */

package service

import (
	"database/sql"
	"exam/Three/dao"
	"exam/Three/model"
)

//注册
func Register(user model.User)error {
	err:=dao.InsertUser(user)
	return err
}

func IsPasswordCorrect(username,password string)(bool,error)  {
	user,err:=dao.SelectUserByUsername(username)
	if err!=nil{
		if err==sql.ErrNoRows{
			return false,nil
		}
		return false, err
	}
	if user.Password!=password{
		return false, err
	}
	return true, err
}

//判断用户名是否存在
func IsExistUsername(username string)(bool,error)  {
	_,err:=dao.SelectUserByUsername(username)
	if err!=nil{
		if err==sql.ErrNoRows{
			return false,nil
		}
		return false, err
	}
	return true,nil
}

func ChangePassword(username, newPassword string) error {
	err:=dao.UpdatePassword(username,newPassword)
	return err
}

//验证密码是否合理(可增加密码复杂性)
func IsPasswordReasonable(password string)bool  {
	if len(password)<6{
		return false
	}
	return true
}

//更新某用户余额
func UpdateMoney(username string,changeMoney int)error  {
	user,err:=dao.SelectUserByUsername(username)
	money:=user.Money+changeMoney
	if money<0{
		return err
	}else {
		err:=dao.UpdateMoney(username,money)
		return err
	}
}

