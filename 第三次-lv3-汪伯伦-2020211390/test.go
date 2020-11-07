package main

import (
	"fmt"
	"os"
)
func main() {
	file, _ := os.OpenFile("./proverb.txt", os.O_RDWR, 0)
	defer file.Close()
	var r [256]byte
	file.Read(r[:])
	fmt.Println(string(r[:]))
}