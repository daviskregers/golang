package main

import (
    "fmt"
)

func sayMessage(msg string, idx int) {
    fmt.Println(msg, idx)
    msg = "123"
}

func sayGreeting(greeting *string, name *string) {
    fmt.Println(*greeting, *name) // Hello Stacey
    *name = "Ted"
    fmt.Println(*name) // Ted
}

func sum(values ...int) (result int) {
    fmt.Println(values)
    for _, v := range values {
        result += v
    }
    return
}

func divide(a, b float64) (float64, error) {
    if b == 0.0 {
        return 0.0, fmt.Errorf("Cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    var message string = "Hello Go!"
    for i := 0; i < 5; i++ {
        sayMessage(message, i)
    }
    fmt.Println(message) // Hello Go!, not 123

    // with pointers

    greeting := "Hello"
    name := "Stacey"
    sayGreeting(&greeting, &name)
    fmt.Println(name) // Ted

    // variatic parameters

    s := sum(1, 2, 3, 4, 5)
    fmt.Println("The sum is ", s)

    // 

    d, err := divide(5.0, 0.0)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(d) 

    // anonymous functions

    func() {
       fmt.Println("Hello Anonymous Go!")
    }() // these () invoke the function

    var f func() = func() {
        fmt.Println("Hello Go!")
    }
    f()

    var divide func(float64, float64) (float64, error)
    divide = func(a, b float64) (float64, error) {
        // ...
        return 0, nil
    }
    divide(0.0, 0.0)

    // methods

    g := greeter {
        greeting: "Hello",
        name: "Go",
    }
    g.greet()
    fmt.Println("The new name is", g.name)
    g.greets()
    fmt.Println("The new name is", g.name)

}

type greeter struct {
    greeting string
    name string
}

func (g greeter) greet() {
    fmt.Println(g.greeting, g.name)
    g.name = "not Go"
}

func (g *greeter) greets() {
    fmt.Println(g.greeting, g.name)
    g.name = "not Go"
}
