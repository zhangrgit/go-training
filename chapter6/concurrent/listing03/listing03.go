/**

 这个示例程序展示如何创建goroutine 以及调度器的行为
 你会看到 goroutine 是并行运行的。两个 goroutine 几乎 是同时开始运行的，大小写字母是混合在一起显示的。
 这是在一台 4核的电脑上运行程序的输出， 所以每个 goroutine 独自运行在自己的核上。
 记住，只有在有多个逻辑处理器且可以同时让每个 goroutine 运行在一个可用的物理处理器上的时候，goroutine 才会并行运行
 */
package  main

import (
	"runtime"
	"sync"
	"fmt"
)

// main是所有Go程序的入口
func main() {

	// 分配2个逻辑处理器给调度器使用
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(2)

	// wg用来等待程序完成
	// 计数加2，表示要等待两个goroutine
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("start goroutines")

	// 声明一个匿名函数，并创建一个goroutine
	go func() {

		// 在函数退出时调用Done来通知main函数工作已经完成
		defer  wg.Done()
		// 显示字母表3次
		for count:=0;count<3;count++{
			for char:='a';char<'a'+26 ;char++  {
				fmt.Printf("%c ",char)
			}
		}
	}()
	// 声明一个匿名函数，并创建一个goroutine
	go func() {
		// 在函数退出时调用Done来通知main函数工作已经完成
		defer  wg.Done()
		// 显示字母表3次
		for count:=0;count<3;count++{
			for char:='A';char<'A'+26 ;char++  {
				fmt.Printf("%c ",char)
			}
		}
	}()
	// 等待goroutine结束
	fmt.Println("waiting to finish")

	wg.Wait()

	fmt.Println("\nTerminating Program")
}