/*******
* @Author:qingmeng
* @Description:
* @File:main
* @Date2021/12/10
 */

package main

import (
	"exam/Three/api"
	"exam/Three/dao"
)

func main() {
	dao.InitDB()
	api.InitEngine()
}
