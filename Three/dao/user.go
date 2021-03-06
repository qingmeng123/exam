/*******
* @Author:qingmeng
* @Description:
* @File:user
* @Date2021/12/10
 */

package dao

import "exam/Three/model"

func InsertUser(user model.User)error  {
	_,err:=DB.Exec("insert into ThirdCoin.user(username, password)values(?,?)",user.Username,user.Password)
	return err
}

func SelectUserByUsername(username string) (model.User, error) {
	user:=model.User{}
	row:=DB.QueryRow("select user_id,password,security_answer from ThirdCoin.user where username=?",username)
	if row.Err()!=nil{
		return user,row.Err()
	}
	err:=row.Scan(&user.UserId,&user.Password,&user.SecurityAnswer)
	if err!=nil{
		return user,err
	}
	return user, err
}

func UpdatePassword(username,newPassword string)error  {
	_,err:=DB.Exec("update ThirdCoin.user set password=? where username=?",newPassword,username)
	return err
}

