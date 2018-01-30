package main

import "fmt"
import "math"

const s string = "const"

func main() {
	fmt.Println(s)
	const n= 5000
	fmt.Println(n)
	const st = "142345"
	fmt.Println(st)

	const d=3e20/n
	fmt.Println(d)
	
	fmt.Println(int64(n))

	fmt.Println(math.Sin(n))
}