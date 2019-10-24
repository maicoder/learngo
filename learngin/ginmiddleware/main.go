package main

import (
	"imooc.com/ccmouse/learngo/learngin/ginmiddleware/initRouter"
)

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()
}
