package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex *sync.RWMutex
	mutex = new(sync.RWMutex)
	fmt.Println("Lock the lock")
	mutex.Lock()
	fmt.Println("The lock is locked")

	channels := make([]chan int, 4)
	for i := 0; i < 4; i++ {
		channels[i] = make(chan int)
		go func(i int, c chan int) {
			fmt.Println("Not read lock: ", i)
			mutex.RLock()
			fmt.Println("Read Locked: ", i)
			fmt.Println("Unlock the read lock: ", i)
			time.Sleep(time.Second)
			mutex.RUnlock()
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

	// Unlock() 使用之前不存在 Lock()
	var rwmutex *sync.RWMutex
	rwmutex = new(sync.RWMutex)
	rwmutex.Unlock()

	// RWMutex 使用不当导致的死锁
	var rwmutexDeadLock *sync.RWMutex
	rwmutexDeadLock = new(sync.RWMutex)
	rwmutexDeadLock.Lock()
	rwmutexDeadLock.Lock()

	//RWMutex 使用不当导致的死锁
	var rwmutexlockRlock *sync.RWMutex
	rwmutexlockRlock = new(sync.RWMutex)
	rwmutexlockRlock.Lock()
	rwmutexlockRlock.RLock()

	//RUnlock() 之前不存在 RLock()
	var rwmutexunLock *sync.RWMutex
	rwmutexunLock = new(sync.RWMutex)
	rwmutexunLock.RUnlock()

	//RUnlock() 个数多于 RLock()
	var rwmutexunlock *sync.RWMutex
	rwmutexunlock = new(sync.RWMutex)
	rwmutexunlock.RLock()
	rwmutexunlock.RLock()
	rwmutexunlock.RUnlock()
	rwmutexunlock.RUnlock()
	rwmutexunlock.RUnlock()

}
