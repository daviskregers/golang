package main

import (
    "fmt"
)

func main() {

    var a int  = 42
    var b *int = &a
    a = 27
    fmt.Println(a) // 27
    fmt.Println(&a) // 0xc0000180d8 (reference)
    fmt.Println(b)  // 0xc0000180d8
    fmt.Println(a)  // 27
    fmt.Println(*b) // 27 (dereference)

    *b = 14
    fmt.Println(a, *b) // 14 14

    c := [3]int{1,2,3}
    d := &c[0]
    e := &c[1] 
    fmt.Println("%v %p %p\n", c, d, e)

    // for pointer arithmetic, we can lookup `unsafe` package

    var ms *myStruct
    ms = &myStruct{foo: 42}
    fmt.Println(ms)

    var mz *myStruct
    fmt.Println(mz) // nil
    mz = new(myStruct)
    fmt.Println(mz) // 0

    var mx *myStruct
    mx = new(myStruct)
    (*mx).foo = 42
    fmt.Println((*mx).foo) // 42
    fmt.Println(mx.foo) // 42

}

type myStruct struct {
    foo int
}
