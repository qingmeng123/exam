/*******
* @Author:qingmeng
* @Description:
* @File:b
* @Date2021/12/26
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	var channel = make(chan int)
	go func() {
		channel<-1
		fmt.Println("下山的路又堵起了")
	}()
	<-channel
	time.Sleep(1 * time.Second)
}