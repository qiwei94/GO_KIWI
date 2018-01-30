package main

import "fmt"

import "time"

func main() {
	i:=2
	fmt.Println("write",i,"as")
	switch i {
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		case 3:
			fmt.Println("three")
	}

	switch time.Now().Weekday(){
		case time.Saturday,time.Sunday:
			fmt.Println("it is weekend")
		default:
			fmt.Println("it is a Weekday")
	}

	/*
	不带表达式的 switch 是实现 if/else 逻辑的另一种方式。这里展示了 case 表达式是如何使用非常量的。
	*/
	t:=time.Now()
	switch{
		case t.Hour()<12:
			fmt.Println("it is before noon")
		default:
			fmt.Println("it is after noon")
	}
}