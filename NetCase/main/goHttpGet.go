package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(resp.ContentLength)
	}