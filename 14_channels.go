package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{}) // signal only channel

func main() {
    ch := make(chan int)

    // Bidirectional comms

    wg.Add(2)
    go func() {
        i := <- ch
        fmt.Println(i) // 42
        ch <- 27
        wg.Done()
    }()
    go func() {
        i := 42
        ch <- i
        fmt.Println(<-ch)
        wg.Done()
    }()
    wg.Wait()

    // 1 send, 1 read thread

    wg.Add(2)

    go func(ch <-chan int) {
        i := <- ch
        fmt.Println(i)
        // ch <- 27 -- cannot send, receive only
        wg.Done()
    }(ch)

    go func(ch chan<- int) {
        ch <- 42
        // fmt.Println(<-ch) -- cannot receive, send only
        wg.Done()
    }(ch)
    wg.Wait()

    // multiple messages

    ch2 := make(chan int, 50)
    wg.Add(2)
    go func(ch <- chan int) {
        for {
            if i, ok := <- ch; ok {
                fmt.Println(i)
            } else {
                break
            }
        }
        wg.Done()
    }(ch2)
    go func(ch chan<- int) {
        ch <- 42
        ch <- 27
        ch <- 13
        close(ch)
        wg.Done()
    }(ch2)
    wg.Wait()

    // logger

    go logger()
    logCh <- logEntry{time.Now(), logInfo, "App is starting"}
    logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
    time.Sleep(100 * time.Millisecond)

    doneCh <- struct{}{}
}

const (
    logInfo = "INFO"
    logWarning = "WARNING"
    logError = "ERROR"
)

type logEntry struct {
    time time.Time
    severity string
    message string
}

func logger() {
    for {
        select {
        case entry := <- logCh:
            fmt.Printf("%v - [%v] %v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
        case <-doneCh:
            break
        }
    }
}
