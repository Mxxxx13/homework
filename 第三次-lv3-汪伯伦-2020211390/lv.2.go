package main

import (
	"os"
)

func main(){
	os.Create ("proverb.txt")
	file,_ := os.OpenFile("./proverb.txt",os.O_WRONLY,0)
	defer file.Close()
	file.WriteString("don't communicate by sharing memory share memory by communicating")
}