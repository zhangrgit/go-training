
/**
  这个示例程序展示如何使用atomic包来提供对数值类型的安全访问
 */
package main

import (
	"sync"
	"sync/atomic"
	"runtime"
	"fmt"
)

var (
	//counter是所有goroutine都要增加其值的变量
	counter int64
	// wg用来等待线程结束
	wg sync.WaitGroup
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

func incCounter(id int)  {

	// 在函数退出时调用done通知main函数工作已完成
	defer  wg.Done()

	for count:=0;count<2 ; count++  {

		// 安全的对counter加1
		atomic.AddInt64(&counter,1)
        // 当前goroutine从线程退出，并放回到队列中
		runtime.Gosched()
	}

}

//Final Counter: 4
// 现在，程序中使用了atmoic包的AddInt64函数。这个函数会同步整型值的加法， 方法是强制同一时刻只能有一个 goroutine 运行并完成这个加法操作。
// 当 goroutine 试图去调用任 何原子函数时，这些 goroutine 都会自动根据所引用的变量做同步处理。现在我们得到了正确的值。


//另外两个有用的原子函数是LoadInt64和StoreInt64。这两个函数提供了一种安全地读 和写一个整型值的方式。
// 代码清单listing05_2.go中的示例程序使用LoadInt64和StoreInt64来创建 一个同步标志，这个标志可以向程序里多个 goroutine 通知某个特殊状态。
