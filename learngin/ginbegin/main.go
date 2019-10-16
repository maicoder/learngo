package main

import (
	"imooc.com/ccmouse/learngo/learngin/ginbegin/initRouter"
)

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()
}
