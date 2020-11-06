package main

import "fmt"

var ch1 = make(chan bool,1)
var ch2 = make(chan bool)
var c = make(chan bool,1)

func GetEvenNum(){
	for i:=0;i<=100;i+=2{
		<- ch1
		fmt.Println(i)
		ch2 <- true
	}
}

func GetOddNum(){
	for i:=1;i<=100;i+=2{
		<- ch2
		fmt.Println(i)
		ch1 <- true
	}
	c <- true
}

func main(){
	ch1 <- true
	go GetEvenNum()
	go GetOddNum()
	<- c
}