### 两种锁
 1. golang 中的 sync 包实现了两种锁。
    - Mutex: 互斥锁
    - RWMutex: 读写锁, RWMutex基于Mutex实现。 
    
 2. MUtex (互斥锁)
    - Mutex 为互斥锁,Lock()加锁，Unlock()解锁
    - 在一个goroutine 获得Mutex 后,其他goroutine 只能等到这个goroutine 释放该Mutex
    - 使用Lock（）加锁后,不能再继续对其加锁,直到Unlock()解锁后才能再加锁.
    - 在lock()之前使用Unlock()解锁会导致panic
    - 已经锁定的Mutex并不与特定的goroutine相关联，这样可以利用一个goroutine对其加锁，再利用其他goroutine对其解锁
    -  在同一个goroutine中的Mutex解锁之前再进行加锁，会导致死锁.
    - 适用于读写不确定，并且只有一个读或者写的场景.
    
  3. RWMutex （读写锁）
    -  RWMtuex 是单写多读锁，该锁可以加多个读锁或者一个写锁.
    - 读锁占用的情况下会阻止写，不会阻止读，多个goroutine可以同时获取读锁
    - 写锁会阻止其他goroutine(无论读写)进来，整个锁由该goroutine 独占
    - 适用于读多写少的场景
    
  4. Lock() 和 Unlock()
   - Lock() 加写锁，Unlock()解写锁
   - 如果在加写锁之前已经有其他读写锁，则Lock()会阻塞直到该锁可用，为确保该锁可用，已经阻塞的Lock()调用会从获得的锁中排除新的读取器，即写锁权限高于读锁，有写锁时优先进行写锁定
   - 在Lock()之前使用Unlock()会导致panic异常
   
   