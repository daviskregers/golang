package main

import (
	"fmt"
	"strconv"
)

// declare at package level
var l float32 = 56

// declare variable block at package level
var (
	actorName 	string = "Elisabeth Sladen"
	companion 	string = "Sarah Jane Smith"
	doctorNumber 	int    = 3
	season 		int    = 11
)

func main() {

	// when we need to declare, but not initialize
	var i int
	i = 42
	i = 27
	fmt.Println(i)

	// when we need to control what type the variable is 
	var j float32 = 42
	fmt.Println(j)

	// 
	k := 12
	fmt.Println(k)

	fmt.Println(l)


	fmt.Println(actorName)

	// redeclare variables

	fmt.Println(l)
	var l int = 42
	fmt.Println(l)

	// variables always have to be used
	// var m int; -- thorws an error "m declared and not used"

	// conversion

	j = float32(l)
	fmt.Printf("%v, %T\n", j, j)

	l = int(j)
	fmt.Printf("%v, %T\n", l, l)

	var m int = 42
	fmt.Printf("%v, %T\n", m, m)

	var n string
	n = string(m)
	fmt.Printf("%v, %T\n", n, n) // will result in `*`

	n = strconv.Itoa(m)
	fmt.Printf("%v, %T\n", n, n) // will result to `42`

}
