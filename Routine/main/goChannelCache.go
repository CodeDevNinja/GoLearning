package main

import "fmt"

func main() {
	//声明一个可以存储整数类型的带缓冲通道
	ch := make(chan int, 2)
	// 因为 ch 是带缓冲的通道,我们可以同时发送两个数据
	// 而不用立刻需要去同步读取数据
	fmt.Println(cap(ch))
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
