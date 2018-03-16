package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"path/filepath"
	"flag"
)

func dirents(dir string)  []os.FileInfo {
	entries,err := ioutil.ReadDir(dir)
	if err != nil{
		fmt.Fprintf(os.Stderr,"du1:%v\n",err)
		return nil
	}
	return entries
}

func walkDir(dir string,fileSizes chan<- int64)  {
	for _,entry := range dirents(dir){
		if entry.IsDir(){
			subdir := filepath.Join(dir,entry.Name())
			walkDir(subdir,fileSizes)
		}else{
			fileSizes <- entry.Size()
		}
	}
}

func main()  {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0{
		roots = []string{"."}
	}
	filesSize := make(chan int64)
	go func() {
		for _,root := range roots{
			walkDir(root,filesSize)
		}
		close(filesSize)
	}()
	var nfiles,nbytes int64
	for size := range filesSize{
		nfiles++
		nbytes += size
	}
	printDiskUsage(nfiles,nbytes)
}
func printDiskUsage(nfiles,nbytes int64)  {
	fmt.Printf("%d files %.1f MB\n",nfiles,float64(nbytes)/1e6)
}
