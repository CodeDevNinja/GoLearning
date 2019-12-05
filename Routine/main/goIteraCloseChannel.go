package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// range 函数遍历每个通道接收到数据， 因为c 在发送完10个
	// 数据之后就关闭了通道， 所以我们range 函数在接收到10个数据
	// 之后就结束了， 如果上面的c通道 不关闭，那么range 函数就不会结束
	//从而接收到第11个 数据的时候就阻塞了
	for i := range c {

		fmt.Println(i)
	}
}
