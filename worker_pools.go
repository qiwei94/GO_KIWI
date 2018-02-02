package main
import "fmt"
import "time"

func worker(id int,jobs <-chan int,results chan<-int) {
	for j:=range jobs{
		fmt.Println("worker",id,"processing job",j)
		time.Sleep(time.Second)
		return <-j *2
	}
}