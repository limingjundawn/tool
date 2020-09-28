package excel

import (
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"io"
	"os"
	"time"
)

func read(p string) string {
	f, err := excelize.OpenFile(p)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")

	//b, err := json.Marshal(rows)
	//fmt.Println("b:", string(b))
	m := [100000]map[string]string{}

	head := map[int]string{}
	for rk, row := range rows {
		tmp := make(map[string]string)
		for ck, colCell := range row {
			if rk == 0 {
				head[ck] = colCell
			} else {
				tmp[head[ck]] = colCell
			}
		}
		if rk == 0 {
			continue
		}
		m[rk-1] = tmp
	}
	j, _ := json.Marshal(m)
	return string(j)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func exists(path string) bool {
	s, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) && !s.IsDir() {
			return true
		}
		return false
	}
	if s.IsDir() {
		return false
	}
	return true
}

func createFlie(filename string, wireteString string) {
	var f *os.File
	var err1 error

	f, err1 = os.Create(filename) //创建文件
	check(err1)
	n, err1 := io.WriteString(f, wireteString) //写入文件(字符串)
	check(err1)
	fmt.Printf("写入 %d 个字节n %s", n, filename)
}

func Change(p string) {
	//查看文件是否存在
	if exists(p) == false {
		fmt.Println("文件不存在")
		os.Exit(0)
	}

	get := read(p)
	time := time.Now().Format("2006-01-02")
	fmt.Println(time)
	createFlie("change_"+string(time)+".json", get)
}
