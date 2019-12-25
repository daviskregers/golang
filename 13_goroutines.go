package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
    for i := 0; i < 10; i++ {
        wg.Add(2)
        go sayHello()
        go increment()
    }
    wg.Wait()

   // 
   fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))
   for i := 0; i < 10; i++ {
       wg.Add(2)
       m.RLock()
       go sayHelloM()
       m.Lock()
       go incrementM()
   }
   wg.Wait()

}

func sayHello() {
    fmt.Printf("Hello #%v\n", counter)
    wg.Done()
}

func sayHelloM() {
    fmt.Printf("Hello #%v\n", counter)
    m.RUnlock()
    wg.Done()
}

func increment() {
    counter++
    wg.Done()
}

func incrementM() {
    counter++
    m.Unlock()
    wg.Done()
}
