```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m sync.Mutex
	cond := sync.NewCond(&m)

	go func() {
		cond.L.Lock()
		fmt.Println("Worker 1: waiting")
		cond.Wait()
		fmt.Println("Worker 1: woke up")
		cond.L.Unlock()
	}()

	go func() {
		cond.L.Lock()
		fmt.Println("Worker 2: waiting")
		cond.Wait()
		fmt.Println("Worker 2: woke up")
		cond.L.Unlock()
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Main: broadcasting")
	cond.Broadcast()

	time.Sleep(1 * time.Second)
	fmt.Println("Main: done")
}
```

| Step | What Happens | Program State |
|------|---------------|----------------|
| **1** | Main starts two goroutines. | Both workers are running concurrently. |
| **2** | Worker 1 calls `Lock()` → gets the lock (nobody else has it). | Lock is **held by Worker 1**. |
| **3** | Worker 1 prints "waiting", then calls `Wait()`. | `Wait()` automatically **unlocks** the lock and **puts Worker 1 to sleep**. |
| **4** | Worker 2 runs, calls `Lock()` → succeeds (lock is free). | Lock is **held by Worker 2**. |
| **5** | Worker 2 prints "waiting", then calls `Wait()`. | `Wait()` unlocks the lock and **puts Worker 2 to sleep**. Both are now sleeping. |
| **6** | After 1 second, main prints "broadcasting" and calls `Broadcast()`. | Both sleeping workers are **woken up**, but they still need to **reacquire the lock**. |
| **7** | One worker (say, Worker 1) reacquires the lock first (maybe), prints "woke up", then unlocks. | Worker 1 finished; Worker 2 waiting for lock. |
| **8** | Worker 2 now gets the lock (maybe), prints "woke up", and unlocks. | Worker 2 finished. |
| **9** | Main prints "done". | All goroutines complete ✅ |
