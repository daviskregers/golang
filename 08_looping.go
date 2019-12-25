package main

import "fmt"

func main() {

    // for loops
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    for i,j := 0, 0; i < 5; i, j = i+1, j+2 {
        fmt.Println(i, j)
    }

    i := 0
    for ; i < 5; i++ {
        // ...
    }
    fmt.Println(i)

    // break

    for {
        fmt.Println(i)
        i++
        if i == 10 {
            break
        }
    }

    // continue

    for i := 0; i < 10; i++ {
        if i % 2 == 0 {
            continue
        }
        fmt.Println(i)
    }

    //

    Loop:
    for i := 1; i <= 3; i++ {
        for j := 1; j <= 3; j ++ {
            fmt.Println(i * j)
            if i * j >= 3 {
                break Loop // break out of both loops at once
            }
        }
    }

    // collections with loops

    s := []int{1,2,3}
    for key, value := range s {
        fmt.Println(key, value)
    }


	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}

    for k, v := range statePopulations {
        fmt.Println(k, v)
    }

    for k := range statePopulations {
        fmt.Println(k)
    }

    for _, v := range statePopulations {
        fmt.Println(v)
    }

    st := "Hello Go!"
    for k, v := range st {
        fmt.Println(k, v, string(v))
    }

}
