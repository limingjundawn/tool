package main

import (
	"fmt"
	"github.com/limingjundawn/tool/excel"
	"os"
)

func main() {
	//读取文件excel路径
	if len(os.Args) <= 1 {
		fmt.Println("请输入文件路径")
		os.Exit(0)
	}

	//只测试了xlsx格式
	excel.Change(os.Args[1])
}
