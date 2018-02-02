package main

import "fmt"
import "time"

func main() {
	queue := make(chan string,2)
	queue<-"one"
	queue<-"two"
	//close(queue)

	go func() {
		for elem:=range queue{
			fmt.Println(elem)
		}
	}()

	for i:=0;i<10;i++{
		fmt.Println(i,":")
		queue<-"i don't know how to transfer int to string"
		time.Sleep(time.Second)
	}

	close(queue)
}