package main

import (
	"imooc.com/ccmouse/learngo/learngin/gintmpl/initRouter"
)

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()
}
