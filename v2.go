package main

import (
	"fmt"
)

var myRes = make(map[int]int, 20)

var channel = make(chan string,200)

func main() {
	for i := 1; i <= 20; i++ {
		go factorial(i)
	}
	for i, v := range myRes {
		fmt.Printf("myres[%d] = %d\n", i, v)
	}
}
func factorial(n int) {
	var res = 1
	for i := 1; i <= n; i++ {
		res *= i
		channel <- string(res)
	}
	myRes[n] = res
}