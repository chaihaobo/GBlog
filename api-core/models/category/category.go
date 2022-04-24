package models

import "github.com/astaxie/beego/orm"

type Category struct {
	Id   int
	Name string `orm:"size(100)"`
}

func insert(category *Category) (bool, error) {
	o := orm.NewOrm()
	i, err := o.Insert(category)
	return i > 0, err
}

func list() []*Category {
	var categorys []*Category
	o := orm.NewOrm()
	o.Read(&categorys)
	return categorys
}

func init() {
	// register model
	orm.RegisterModel(new(Category))
}
