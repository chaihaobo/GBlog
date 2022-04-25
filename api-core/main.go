package main

import (
	_ "api-core/models/category"
	_ "api-core/routers"
	"github.com/astaxie/beego/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "./apicore.db")
	orm.RunSyncdb("default", false, true)
}

func main() {
	orm.Debug = true
	beego.Run()
}
