/*******
* @Author:qingmeng
* @Description:
* @File:middleware
* @Date2021/12/10
 */

package api

import (
	"exam/Three/tool"
	"github.com/gin-gonic/gin"

)

func auth(ctx *gin.Context) {
	username,err:=ctx.Cookie("username")
	if err!=nil{
		tool.RespErrorWithData(ctx,"请先登录")
		ctx.Abort()
	}
	ctx.Set("iUsername",username)
	ctx.Next()
}
