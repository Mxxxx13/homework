package main

import "fmt"

func Receiver(v interface{}){
	switch v.(type){
	case int:
		fmt.Println("这个是int")
	case string:
		fmt.Println("这个是string")
	case bool:
		fmt.Println("这个是bool")
	}
}

func main() {
	v1 := 13
	v2 := true
	v3 := "hello world"
	Receiver(v1)
	Receiver(v2)
	Receiver(v3)
}
