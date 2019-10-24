package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	for inx, arg := range os.Args {
		fmt.Println("参数"+strconv.Itoa(inx)+":", arg)
	}

	path := RootPath()
	fmt.Println(path)
}

func RootPath() string {
	s, err := exec.LookPath(os.Args[0])
	if err != nil {
		log.Panicln("发生错误", err.Error())
	}
	i := strings.LastIndex(s, "/")
	path := s[0: i+1]
	return path
}