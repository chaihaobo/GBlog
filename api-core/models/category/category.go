package category

import "github.com/astaxie/beego/orm"
import "api-core/models"

func Insert(category *models.Category) (bool, error) {
	o := orm.NewOrm()
	i, err := o.Insert(category)
	return i > 0, err
}

func List() []*models.Category {
	var categorys []*models.Category
	o := orm.NewOrm()
	o.QueryTable(new(models.Category)).All(&categorys)
	return categorys
}
