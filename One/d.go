/*******
* @Author:qingmeng
* @Description:
* @File:d
* @Date2021/12/26
 */



package main

import (
	"fmt"
)

//向 numChan放入 1000000个数
func inputNum(numChan chan int) {
	//放入1000000个数
	for i := 1; i < 1000000; i++ {
		numChan <- i
	}
	close(numChan)

}

//将numChan中的数取出，找出素数并放入primeNumChan
func Prime(numChan, primeNumChan chan int, exitChan chan bool) {
	//取出管道中的值
	for {
		num, ok := <-numChan
		flag := true
		//管道中没有值时结束
		if !ok {
			break
		}
		//判断素数
		for j := 2; j < num; j++ {
			if num%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeNumChan <- num
		}
	}
	exitChan <- true
}

func main() {
	//带有10000个缓存的numChan
	numChan := make(chan int, 10000)
	primeNumChan := make(chan int, 2000)
	exitChan := make(chan bool, 10)

	go inputNum(numChan)
	//开启10个协程进行方法prime
	for i := 1; i <= 10; i++ {
		go Prime(numChan, primeNumChan, exitChan)
	}
	//处理协程
	go func() {
		for i := 1; i <= 10; i++ {
			<-exitChan
		}
		close(primeNumChan)
	}()
	//遍历输出结果
	for {
		prime, ok := <-primeNumChan
		if !ok {
			break
		}
		fmt.Println(prime)
	}

}
