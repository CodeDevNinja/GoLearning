package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	fmt.Println("Lock the lock")
	mutex.Lock()
	fmt.Println("The lock is locked")
	channels := make([]chan int, 4)
	for i := 0; i < 4; i++ {
		channels[i] = make(chan int)
		go func(i int, c chan int) {
			fmt.Println("Not lock: ", i)
			mutex.Lock()
			fmt.Println("Locked: ", i)
			time.Sleep(time.Second)
			fmt.Println("Unlock the lock: ", i)
			mutex.Unlock()
			c <- i
		}(i, channels[i])
	}
	time.Sleep(time.Second)
	fmt.Println("Unlock the lock")
	mutex.Unlock()
	time.Sleep(time.Second)

	for _, c := range channels {
		<-c
	}


//	解锁之前会加锁导致死锁
//var mutexDeadlock sync.Mutex
//	mutexDeadlock.Lock()
//	fmt.Println("Locked")
//	mutexDeadlock.Lock()

// 在 Lock() 之前使用 Unlock() 会导致 panic 异常
	var mutexPanic sync.Mutex
	mutexPanic.Unlock()


}
