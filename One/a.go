/*******
* @Author:qingmeng
* @Description:
* @File:a
* @Date2021/12/26
 */

package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex

	go func() {
		fmt.Println("有点强人锁男")
		mu.Lock()
	}()

	mu.Unlock()
}
//主协程比子协程执行的快，主协程直接运行了unlock,而分协程还没执行lock.协程报错