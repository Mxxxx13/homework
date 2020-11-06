package main

import (
	"fmt"
)

var (
	myres = make(map[int]int, 20)
	ch = make(chan int)
	n int
)

func factorial(n int) {
	var res = 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	myres[n] = res
	ch <- res
}

func main() {
	for i := 1; i <= 20; i++ {
		go factorial(i)
	}
	myres[n]= <- ch
	for i, v := range myres {
		fmt.Printf("myres[%d] = %d\n", i, v)
	}

}
