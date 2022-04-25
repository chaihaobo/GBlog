package routers

import (
	"api-core/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	categoryNs := beego.NewNamespace("/category",
		beego.NSRouter("/list", &controllers.CategoryController{}, "get:List"),
		beego.NSRouter("/save", &controllers.CategoryController{}, "post:Save"),
		beego.NSRouter("/:id:int", &controllers.CategoryController{}, "delete:Delete"),
	)
	beego.AddNamespace(categoryNs)
}
