// 这个示例程序展示如何在程序里造成竞争状态,实际上不希望出现这种情况
// go build -race   // 用竞争检测器标志来编译程序   ./example
package main

import (
	"sync"
	"runtime"
	"fmt"
)

var(
	// counter是所有goroutine都要增加其值的变量
	counter int
	// wg用来等待程序结束
	wg sync.WaitGroup
)

// main是所有Go程序的入口
func main() {

	// 计数加2，表示要等待两个goroutine
	wg.Add(2)

	// 创建两个goroutine
	go incCounter(1)
	go incCounter(2)

	// 等待goroutine结束
	wg.Wait()

	fmt.Println("final counter:",counter)
}
// incCounter增加包里counter变量的值
func incCounter(id int)  {
	// 在函数退出时调用Done来通知main函数工作已经完成
	defer  wg.Done()
	for count := 0; count < 2; count++ {
		// 捕获counter的值
		value:=counter
		// 当前goroutine从线程退出，并放回到队列
		runtime.Gosched()
		// 增加本地value变量的值
		value++
		// 将该值保存回counter
		counter=value
	}
}
