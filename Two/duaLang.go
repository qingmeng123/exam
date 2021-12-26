/*******
* @Author:qingmeng
* @Description:
* @File:duaLang
* @Date2021/12/26
 */

package main

import "fmt"

type room struct {
	stair	int		//1为有楼梯，0为无
	number  int 	//编号，从0开始
	indicator int 	//指示牌，表示从这个房间开始按逆时针方向选择第x个有楼梯的房间。
}

type floor struct {
	number 	int 		//楼层数
	rooms []room		//每层楼房间
}

type house struct {
	floor []floor
}

func createRoom(stair,number,indicator int)room  {
	return room{
		stair:     stair,
		number:    number,
		indicator: indicator,
	}
}

func createFloor(number int ,rooms []room)floor  {
	return floor{
		number: number,
		rooms: rooms,
	}
}

func createHouse() house {
	var n,m int
	//输入n个楼层和m个房间
	fmt.Scan(&n,&m)
	rooms:=make([]room,m)
	floors:=make([]floor,n)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var stair,indicator int
			//输入0、1表有无楼梯，输入指示牌
			fmt.Scan(&stair,&indicator)
			rooms[j]=createRoom(stair,j,indicator)
		}
		floors[i]=createFloor(i,rooms)
	}
	return house{floor: floors}
}

func main()  {
	house:=createHouse()
	var intRoom int //从第几个房间进入
	fmt.Scan(&intRoom)
	code:=0	//密码

	for i := 0; i < len(house.floor); i++ {
		code+=house.floor[i].rooms[intRoom].indicator
		key:=0	//记录按逆时针排的第几个有楼梯的房间
		for{
			if house.floor[i].rooms[intRoom].stair==1{
				key++
			}
			if key==house.floor[i].rooms[intRoom].indicator{
				intRoom=house.floor[i].rooms[intRoom].number
				break
			}
			intRoom++
			intRoom=intRoom% len(house.floor[i].rooms)
		}

	}
	fmt.Println(code)
}