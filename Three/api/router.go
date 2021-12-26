/*******
* @Author:qingmeng
* @Description:
* @File:router
* @Date2021/12/10
 */

package api

import "github.com/gin-gonic/gin"

func InitEngine() {
	engine:=gin.Default()

	engine.POST("/user",register)		//注册
	engine.GET("/user",login)			//登陆

	userGroup:=engine.Group("/user")
	{
		userGroup.POST("/password")

		passwordGroup:=userGroup.Group("/password")
		{
			passwordGroup.POST("/",auth,changePassword)	//登陆后的直接修改密码
		}
		userGroup.POST("/updateMoney",auth,UpdateMoney)	//修改某账户余额
	}
	engine.Run()
}
