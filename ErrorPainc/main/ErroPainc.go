package main

import (
	"os"
	"log"
	"syscall"
	"fmt"
	"io"
)

func fatal()  {
		m := make(map[string]string)
	m ["key"]= "123"
	if v, ok := m["key"];ok{
		fmt.Println(v)
	}else{
		fmt.Println("Error")
	}

	err := syscall.Chmod("",0666)
	if err != nil{
		log.Fatal(err.(syscall.Errno))
	}
}
func main()  {
	CopyFile("/home/caidong/goLearning/ErrorPainc/main/m.go","/home/caidong/goLearning/ErrorPainc/main/ErroPainc.go")
}

func CopyFile(destName,srcName string)(written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		log.Fatal(err.(syscall.Errno))
		return
	}
	dst,err := os.Create(destName)
	if err != nil{
		log.Fatal(err.(syscall.Errno))
		return
	}
	defer src.Close()
	return io.Copy(dst, src)
}