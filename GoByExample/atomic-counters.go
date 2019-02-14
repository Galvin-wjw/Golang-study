package main

import (
	"sync/atomic"
	"runtime"
	"time"
	"fmt"
)

func main() {
	var ops uint64 = 0
	for i := 0 ; i < 50; i++{
		go func() {
			for {
				atomic.AddUint64(&ops,1)

				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	//extract a copy of the current value into opsFinal via LoadUint64.
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:",opsFinal)
}
