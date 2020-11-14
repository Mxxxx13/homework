package main

import (
	"fmt"
	"time"
)

var c = make(chan int)

func Fly(){
	timer := time.Tick(time.Hour)
	for _ = range timer{
		fmt.Println("芜湖,起飞")
	}
	c <- 1
}

func main(){
	go Fly()
	<- c
}

//只会芜湖,我是five