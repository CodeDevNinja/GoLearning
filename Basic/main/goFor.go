package main

import "fmt"

func main() {
	for index,i := range "1, 23" {
		fmt.Println(i,index)
	}
}