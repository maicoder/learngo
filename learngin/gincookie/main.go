package main

import (
	"imooc.com/ccmouse/learngo/learngin/gincookie/initRouter"
)

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()
}
