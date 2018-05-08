/*
  从标准输入中读取文本，并进行下面的操作
  1.计算字符数量（包括空格）
  2. 计算单词数量
  3. 计算行数
*/

package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main()  {
	var chars, words,lines int
	r := bufio.NewReader(os.Stdin)
	for{
		switch s,ok := r.ReadString('\n');true{
		case ok != nil:
			fmt.Printf("chars is %d,words is %d,lines is %d",chars,words,lines)
			return
		default:
			chars += len(s)
			words += len(strings.Fields(s))
			lines++
		}
	}
}
