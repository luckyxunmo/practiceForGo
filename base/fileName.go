package main

import (
	"fmt"
	"path"
	"strings"
)

func main() {
	fullFilename := "/Users/itfanr/Documents/test.txt"
	fmt.Println("fullFilename =", fullFilename)
	var filenameWithSuffix string
	filenameWithSuffix = path.Base(fullFilename)            //获取文件名带后缀
	fmt.Println("filenameWithSuffix =", filenameWithSuffix) //filenameWithSuffix = test.txt
	var fileSuffix string
	fileSuffix = path.Ext(filenameWithSuffix) //获取文件后缀
	fmt.Println("fileSuffix =", fileSuffix)   //fileSuffix = .txt

	var filenameOnly string
	filenameOnly = strings.TrimSuffix(filenameWithSuffix, fileSuffix) //获取文件名
	fmt.Println("filenameOnly =", filenameOnly)                       //filenameOnly = test
}
