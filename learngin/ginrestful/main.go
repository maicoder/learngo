package main

import (
	"imooc.com/ccmouse/learngo/learngin/ginrestful/initRouter"
)

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()
}
