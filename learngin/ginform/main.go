package main

import (
	"imooc.com/ccmouse/learngo/learngin/ginform/initRouter"
)

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()
}
