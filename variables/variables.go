package main

import (
	"fmt"
	"strconv"
)

var (
	actorName string = "John Travolta"
	companion string = "Nicolas Cage"
	id        int    = 3
) // can group related variables for code organization

var i int = 47

func main() {
	var i int            //name, type
	i = 42               // initialize variable
	var j float32 = 27   //name, type, value (must be use in package level variables)
	k := 99              // auto detects type (int in this case)
	l := 99.             // auto detects type (float in this case)
	m := float32(k)      // convert from one type to another
	s := strconv.Itoa(i) // convert a int to string using strconv package
	fmt.Println("integer:", i, j, k, l, m, s)
}
