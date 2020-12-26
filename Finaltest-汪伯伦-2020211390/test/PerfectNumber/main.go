package main

import (
	"fmt"
	"time"
)

func main(){
	for i:=1;i<100;i++{
		for j:=0;j<124;j++{
			go PerfectNumber(i+j*100)
		}
	}
	time.Sleep(5)
}

func PerfectNumber(num int){
	FactorSum(num)
}

func FactorSum(num int){
	var sum int
	for i:=1;i<num;i++ {
		if num%i == 0 {
			sum += i
		}
	}
	check(num,sum)
}

func check(num int,sum int)  {
	if num == sum {
		fmt.Printf("perfect number is %d\n", num)
	}
}