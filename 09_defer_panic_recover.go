package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

func main() {

    fmt.Println("start")
    defer fmt.Println("middle") // execute at the end of the function, before return
    fmt.Println("end")

    // start
    // end
    // middle

    defer fmt.Println("start")
    defer fmt.Println("middle")
    defer fmt.Println("end")

    // Executes in LIFO:
    // end
    // middle
    // start

    res, err := http.Get("http://www.google.com/robots.txt")
    defer res.Body.Close()

    if err != nil {
        log.Fatal(err)
    }

    robots, err := ioutil.ReadAll(res.Body)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%s\n", robots[0:100])

    //

    defer func() {
        if err := recover(); err != nil {
            log.Println("Error: ", err)
        }
    }()

    a, b := 1, 0
    ans := a / b
    fmt.Println(ans)

    // will not execute this because 1 / 0 threw a panic, stopped execution
    panic("Something bad happened!")

}
