package main

import (
	"fmt"
	"time"

	"imooc.com/ccmouse/learngo/retriever/mock"
	"imooc.com/ccmouse/learngo/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

//func download(r Retriever) string {
//	return r.Get("http://www.imooc.com")
//}

func main() {
	var r Retriever
	r = mock.Retriever{"this is a fake imooc.com"}
	fmt.Printf("%T %v\n", r, r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		Timeout:   time.Minute,
	}
	fmt.Printf("%T %v\n", r, r)

	inspect(r)

	// Type assertion
	if mockRetriever, ok := r.(mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	//fmt.Println(download(r))
}

func inspect(r Retriever) {
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent: ", v.UserAgent)
	}
}
