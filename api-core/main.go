package main

import (
	_ "api-core/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

