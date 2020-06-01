package main

import (
	_ "Logical_server/System/routers"
	_ "Logical_server/System/sysinit"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
