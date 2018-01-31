package main
import "fmt"


func inSeq() func() int {
	i:=0
	return func() int {
		i+=1
		return i
	}
}

func main() {
	nextInt:=inSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts:=inSeq()
	fmt.Println(newInts())
}