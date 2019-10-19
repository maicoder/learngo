package main

import (
	"imooc.com/ccmouse/learngo/learngin/ginmysql/initRouter"
)

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()
}
