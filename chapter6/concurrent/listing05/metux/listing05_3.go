/**
  这个示例程序展示如何使用互斥锁来定义一段需要同步访问的代码临界区资源的同步访问
 */
package main

import (
	"sync"
	"fmt"
	"runtime"
)

var (
	//counter是所有goroutine都要增加其值的变量
	counter int64
	// wg用来等待线程结束
	wg sync.WaitGroup
   // 用来定义一段代码临界区
	mutex sync.Mutex
)

func main() {

	//计数加2，表示要等待两个goroutine
	wg.Add(2)

	//创建两个goroutine
	go incCounter(1)
	go incCounter(2)

	// 等待goroutine结束
	wg.Wait()

	//显示最终的值
	fmt.Println("Final Counter:",counter)

}

// incCounter使用互斥锁来同步并保证安全访问，增加包里counter变量的值
func incCounter(id int)  {

	// 在函数退出时调用done通知main函数工作已完成
	defer  wg.Done()

	for count:=0;count<2 ; count++  {

		//同一个时刻只允许一个goroutine进入这个临界区
		mutex.Lock()
		{

			//捕获counter的值
			value := counter
			//当前goroutine从线程退出，并释放到队列中
			runtime.Gosched()
			// 增加本地value变量的值
			value++
			//将该值保存回counter
			counter = value

		}
		//释放锁，允许其它正在等待的goroutine进入临界区
		mutex.Unlock()


	}

}

//Final Counter: 4

//对counter变量的操作在Lock()和Unlock()函数调用定义的临界 区里被保护起来。使用大括号只是为了让临界区看起来更清晰，并不是必需的。
// 同一时刻只有一 个 goroutine 可以进入临界区。之后，直到调用Unlock()函数之后，其他 goroutine 才能进入临界区。
// 当强制将当前 goroutine 退出当前线程后，调度器会再次分配这个 goroutine 继续运行。当程序结束时，我们得到正确的值4，竞争状态不再存在。
