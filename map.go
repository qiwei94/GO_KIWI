package main

import "fmt"

func main() {
	m:=make(map[string]int)

	m["k1"]=4
	m["k2"]=5
	fmt.Println("map",m)

	v1:=m["k1"]
	fmt.Println("v1:",v1)

	fmt 

}