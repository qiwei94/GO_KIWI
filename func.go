package main

import "fmt"

func plus(a int, b int) int{
	return a+b
}

func main() {
	res:=plus(2,3)
	fmt.Println("res:",res)
}