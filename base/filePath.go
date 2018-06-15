package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	getAbsPath()
}
func filepathPractise() {

	execpath, _ := os.Executable() // 获得程序路径

	configfile := filepath.Join(filepath.Dir(execpath), "./config.yml") //拼接path
	fmt.Println(execpath, configfile)
	currentDirPath, fileName := filepath.Split(configfile)
	fmt.Println(currentDirPath, "\n", fileName)

}
func getAbsPath() {
	dirPath := "../upload"
	dfi, err := os.Open(dirPath)
	if err != nil {
		fmt.Println("open:", dirPath, "err:", err.Error())
	}
	defer dfi.Close()
	if !filepath.IsAbs(dirPath) {
		absDirPath, err := filepath.Abs(dirPath)
		if err != nil {
			fmt.Errorf("cannot get absolute path of directory: %s", err.Error())
		} else {
			fmt.Println("absdir:", absDirPath)
		}
	} else {
		fmt.Println("it is abs:", dirPath)
		configfile := filepath.Join(dirPath, "config.yml")
		fmt.Println(configfile)
	}

}
func checkDir(dir string) (bool, error) {
	exist, err := PathExists(dir)
	if err != nil {
		//  log.Error("get dir error：", err.Error())
		fmt.Println("get dir error：", err.Error())
		return false, err
	}

	if exist {
		return true, nil
	} else {
		// 创建文件夹
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			//log.Errorf("mkdir failed!dir:%s,error:[%s]\n", dir, err.Error())

			fmt.Println("mkdir failed! error：", err.Error())
			return false, err
		} else {
			return true, nil
		}
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
