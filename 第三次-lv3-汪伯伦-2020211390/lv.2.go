package main

import (
	"fmt"
	"os"
)

func main(){
	//os.Create ("proverb.txt")
	file,_ := os.OpenFile("./proverb.txt",os.O_RDWR,0)
	defer file.Close()
	//file.WriteString("don't communicate by sharing memory share memory by communicating")
	var r [256]byte
	file.Read(r[:])
	fmt.Println(string(r[:]))
}