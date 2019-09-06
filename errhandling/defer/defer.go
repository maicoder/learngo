package main

import "fmt"

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("error occurred")
	fmt.Println(4)
}

func writeFile(filen ame string)  {
	
}

func main() {
	tryDefer()
}
