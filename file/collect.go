package file


import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

//移动除了快捷方式的文件，到统一的文件夹里面

var paths  string//路径

func setPaths(s string){
	paths = s
}


func readFile() []string {
	if paths == ""{
		fmt.Println("未设置路径")
		os.Exit(0)
	}
	files, _ := ioutil.ReadDir(paths)
	s := make([]string, 100)
	for _, f := range files {
		s = append(s, f.Name())
	}
	return  s
}

//移动文件
func moveFile(s []string, p string) {
	for _, v := range s{
		if v == "" {
			continue
		}
		//判断是否是文件夹
		fileInfo, _ := os.Stat(paths + "\\" +v)
		if fileInfo.IsDir() {
			continue
		}
		//判断文件类型，除了快捷方式，统统移动
		if path.Ext(v) != ".lnk"{
			errs := os.Rename(paths + "\\" +v, p + "\\" + v)
			if errs != nil{
				log.Fatal(errs)
			}
		}
	}
}

func pathExists(path string) (bool, error){
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err){
		return false, nil
	}
	return false, err
}

func ClearFile(p string, np string)  {
	res, _ := pathExists(np)
	if !res {
		err := os.Mkdir(np, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	setPaths(p)
	rres := readFile()
	moveFile(rres, np)
}
