package main

import "fmt"

func main() {
	over := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
			<- over 		//对于无缓存的channel接受要在发送前
		}()
		if i <= 9 {			//只要符合条件则打印i， i==9 只能打印出9
			over <- true
		}
	}
	fmt.Println("over!!!")
}