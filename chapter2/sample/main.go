package main

import (
	"log"
	"os"
	_ "github.com/zhangrgit/go-training/chapter2/sample/matchers" //调用 matchers包中的 rss.go 代码文件里的init函数，注册 RSS 匹配器，以便后用
	"github.com/zhangrgit/go-training/chapter2/sample/search"
)

// init在main之前调用
func init() {
	//将日志输出到标准输出
	 log.SetOutput(os.Stdout)
}

func main() {

	// 使用特定的项做搜索
	search.Run("president")

}