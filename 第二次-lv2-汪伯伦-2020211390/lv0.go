package main

import (
	"fmt"
)

type cat struct {
	name string
	color string
	gender string
	age int
}

func main(){
	var c cat
	c.name = "feibo"
	c.color = "yellow"
	c.gender = "male"
	c.age = 3
	fmt.Println(c)
}