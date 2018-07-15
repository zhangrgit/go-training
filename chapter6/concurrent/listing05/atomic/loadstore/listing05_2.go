
/**
   这个示例程序展示如何使用atomic包里的Store和Load类函数来提供对数值类型的安全访问
 */
package main

import (
	"sync"
	"sync/atomic"
	"fmt"
	"time"
)

var (
	//shutdown 是通知正在执行的goroutine停止工作的标志
	shutdown int64
	// wg用来等待线程结束
	wg sync.WaitGroup
)

func main() {

	//计数加2，表示要等待两个goroutine
   wg.Add(2)

   //创建两个goroutine
   go doWork("A")
   go doWork("B")

   // 给定goroutine执行的时间（1秒）
   time.Sleep(1 * time.Second)

   // 该停止工作了，安全的设置shutdown标志
   fmt.Println("Shutdown Now")
   atomic.StoreInt64(&shutdown,1)

   //等待goroutine结束
   wg.Wait()

}

func doWork(name string )  {

	// 在函数退出时调用done通知main函数工作已完成
	defer  wg.Done()

	for {
		fmt.Printf("Doing %s Work\n",name)
		//模拟任务执行时间250毫秒
		time.Sleep(250 * time.Millisecond)

		//要停止工作了吗?
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n",name)
			break

		}

	}

}

//在这个例子中，启动了两个 goroutine，并完成一些工作。在各自循环的每次迭代之后，在代码中 goroutine 会使用 LoadInt64 来检查 shutdown 变量的值。
// 这个函数会安全地返回 shutdown变量的一个副本。如果这个副本的值为 1，goroutine 就会跳出循环并终止。

// 使用StoreInt64函数来安全地修改shutdown变量的值。
// 如果 哪个doWork goroutine 试图在main函数调用StoreInt64的同时调用LoadInt64函数，那 么原子函数会将这些调用互相同步，保证这些操作都是安全的，不会进入竞争状态。
