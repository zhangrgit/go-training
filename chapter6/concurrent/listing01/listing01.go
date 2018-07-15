/**
  创建两个 goroutine，以并发的形式分别显示大写和小写的英文字母。
 这个示例程序展示如何创建goroutine 以及调度器的行为
 */
package  main

import (
	"runtime"
	"sync"
	"fmt"
)

// main是所有Go程序的入口
func main() {

	// 分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(1)

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