package main
import "fmt"
import "time"

func worker(done chan bool){
	fmt.Println("working....")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func  main() {
	done:=make(chan bool,1)
	go worker(done)
	for i:=1;i<10;i++{
		fmt.Println(i)
	}
	<-done
	for j:=11;j<20;j++{
		fmt.Println(j)
	}
}