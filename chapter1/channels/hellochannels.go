package main

import (
	"fmt"
	"sync"
	"time"
)

var  wg sync.WaitGroup
func printer(ch chan  int)  {
	for i:=range ch{
	   fmt.Printf("Received%d ",i)
	}
	println()
	//模拟业务过程
	time.Sleep(1000)
	defer wg.Done()
}
// main是程序的入口点
func main() {

	c:=make(chan int)
	// channel须先接收后往其发送，否则死锁：fatal error: all goroutines are asleep - deadlock!
	go printer(c)
	wg.Add(1)
	//在通道上发送10个整数
	for i:=1;i<=10 ; i++ {
		c <- i
	}
	close(c)
	//等待协程处理结束
	wg.Wait()
	// 理论上finish是在协程完成后最后打印，不过当协程任务时间过短时，协程调度器未能及时的处理这一调度进行等待，最终的顺序可能会发生变化
	println("finish!")



}





