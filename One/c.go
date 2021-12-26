/*******
* @Author:qingmeng
* @Description:
* @File:c
* @Date2021/12/26
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	var channel = make(chan int)
	for i := 0; i < 10; i++ {
		go func() {
			channel<-1
			fmt.Println("这是第",i+1,"句话")
		}()
		<-channel
		time.Sleep(1 * time.Second)
	}

}
