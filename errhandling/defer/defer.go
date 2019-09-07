package main

import (
	"bufio"
	"fmt"
	"imooc.com/ccmouse/learngo/functional/fib"
	"os"
)

func tryDefer() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close() // defer 语句以栈的形式存储，函数结束后，函数里面的defer语句以先进后出的顺序执行

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	tryDefer()
	writeFile("fib.txt")
}
