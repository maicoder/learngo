package main

import (
	"imooc.com/ccmouse/learngo/learngin/ginupload/initRouter"
)

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()
}
