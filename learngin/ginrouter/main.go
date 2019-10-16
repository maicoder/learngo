package main

import (
	"imooc.com/ccmouse/learngo/learngin/ginrouter/initRouter"
)

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()
}
