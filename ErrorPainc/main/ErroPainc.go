package main

import (
	"log"
	"syscall"
	"fmt"
)

func main()  {
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